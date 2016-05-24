package invoke

import (
	"scheduler/entity"
	"time"
	"log"
	"scheduler/policy"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"scheduler/common"
	"bytes"
)

// 执行
type Invoker struct {

}

// 执行任务
func (this *Invoker)Execute(jobInfo *entity.JobInfo,nextTime time.Time,params string) error {
	snapshot, err := this.Init(jobInfo,nextTime)
	log.Println("snapshot = ", snapshot)
	fac := &policy.Factory{}
	policy := fac.FindPolicy(jobInfo)
	count := 0;
	for {

		url := policy.GetNextUrl()

		if url == "" {
			this.executeJobResult(snapshot, "所有目标服务器地址都不可用" + time.Now().Local().Format("2006-01-02 15:04:05"), entity.ERROR)
			break;
		}
		snapshot.Url = url
		// 准备执行

		err = this.invoke(snapshot)
		if err == nil {
			break
		} else if count < 3 {

			this.executeJobResult(snapshot, "目标服务器地址:"+url+"不可用正在尝试" + time.Now().Local().Format("2006-01-02 15:04:05"), entity.ERROR)
		}
		count ++


	}
	return err
}


// 执行任务
func (this *Invoker) invoke(jobSnapshot *entity.JobSnapshot) error {

	this.executeJobResult(jobSnapshot,"目标服务器地址:" + jobSnapshot.Url + "正在执行" + time.Now().Local().Format("2006-01-02 15:04:05"),entity.EXECUTING)
	startTime := time.Now()

	jobRequest := &common.JobRequest{JobSnapshot:jobSnapshot.Id,Params:jobSnapshot.Params,Status:common.EXECUTING}

	content,err := json.Marshal(jobRequest)

	if err != nil {
		this.executeJobResult(jobSnapshot,"解析job请求参数出错" + time.Now().Local().Format("2006-01-02 15:04:05"),entity.ERROR)
		return err
	}

	resp,err := http.Post(jobSnapshot.Url,"application/json;charset=utf-8",bytes.NewBuffer(content))
	if err != nil {
		this.executeJobResult(jobSnapshot,"目标服务器地址:" + jobSnapshot.Url + "不可用" + time.Now().Local().Format("2006-01-02 15:04:05"),entity.ERROR)
		return  err
	}
	defer  resp.Body.Close()
	body,err := ioutil.ReadAll(resp.Body)
	result := &common.JobResponse{}
	log.Println("body = ",string(body))
	err = json.Unmarshal(body,result)

	if err != nil {
		this.executeJobResult(jobSnapshot,"目标服务器地址:" + jobSnapshot.Url + "非法的响应" + time.Now().Local().Format("2006-01-02 15:04:05"),entity.ERROR)
		log.Println("err = ",err)
		return  err
	}

	addr := resp.Request.RemoteAddr

	log.Println("serveraddr=",addr)
	jobSnapshot.ServerAddress = addr
	if result.Success == true {
		nowTime := time.Now()
		d :=nowTime.Sub(startTime)
		timeConsume := int64(d.Seconds())
		jobSnapshot.TimeConsume = timeConsume
		jobSnapshot.Result = result.Content
		jobSnapshot.Ip = common.GetIPFromUrl(jobSnapshot.Url)
		this.executeJobResult(jobSnapshot,"目标服务器地址:" + jobSnapshot.Url + "执行成功 返回值:" + result.Message + time.Now().Local().Format("2006-01-02 15:04:05"),entity.COMPLETED)

	} else {
		this.executeJobResult(jobSnapshot,"目标服务器地址:" + jobSnapshot.Url + "执行失败:" + result.Message + time.Now().Local().Format("2006-01-02 15:04:05"),entity.ERROR)

	}
	return nil

}

func (this *Invoker)executeJobResult(snapshot *entity.JobSnapshot, detail, status string) {

	snapshot.Detail = snapshot.Detail + detail + "\n"
	snapshot.Status = status
	snapshot.ModifyTime = time.Now().Local()
	snapshot.UpdateSnapshot()

}
func (this *Invoker) Init(jobInfo *entity.JobInfo,nextTime time.Time) (*entity.JobSnapshot, error) {

	serverAddr := common.GetLocalAddr()
	snapshot := &entity.JobSnapshot{
		JobInfoId:jobInfo.Id,
		Name          :jobInfo.Name,
		Group         :jobInfo.Group,
		Status        :entity.INIT,
		Url           :jobInfo.Urls,
		TimeConsume   :0,
		ServerAddress:serverAddr,
		NextTime :nextTime,
		CreateTime    :time.Now().Local(),
		Detail:"初始化" + time.Now().Local().Format("2006-01-02 15:04:05") + "\n",
		Params:jobInfo.Param,

	}
	err := snapshot.InsertJobSnapshot()
	return snapshot, err
}

