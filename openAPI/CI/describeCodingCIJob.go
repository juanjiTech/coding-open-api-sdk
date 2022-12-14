package ci

import "github.com/juanjiTech/coding-open-api-sdk/openAPI/poster"

type DescribeCodingCIJobReq struct {
Response struct {
Job struct {
AlwaysUserIDList []interface{} `json:"AlwaysUserIdList"` // <nil>
AutoCancelSameMergeRequest bool `json:"AutoCancelSameMergeRequest"` // true
AutoCancelSameRevision bool `json:"AutoCancelSameRevision"` // true
BranchRegex string `json:"BranchRegex"` // ^refs/heads/master$
BranchSelector string `json:"BranchSelector"` // master
BuildFailUserIDList []interface{} `json:"BuildFailUserIdList"` // <nil>
CachePathList []interface{} `json:"CachePathList"` // <nil>
CreatedAt int64 `json:"CreatedAt"` // 0
CreatorID int64 `json:"CreatorId"` // 1
DepotHttpsURL string `json:"DepotHttpsUrl"` // http://e.coding.9.134.115.58.nip.io/codingcorp/test-1.git
DepotID int64 `json:"DepotId"` // 1
DepotName string `json:"DepotName"` // test-1
DepotSshURL string `json:"DepotSshUrl"` // git@e.coding.9.134.115.58.nip.io:codingcorp/test-1.git
DepotType string `json:"DepotType"` // CODING
DepotWebURL string `json:"DepotWebUrl"` // http://codingcorp.coding.9.134.115.58.nip.io/p/test-1/d/test-1/git
DockerBuildPath string `json:"DockerBuildPath"` // 
DockerBuildTag string `json:"DockerBuildTag"` // 
DockerFilePath string `json:"DockerFilePath"` // 
EnvList []struct {
Name string `json:"Name"` // a
Sensitive bool `json:"Sensitive"` // false
Value int64 `json:"Value,string"` // 1
} `json:"EnvList"`
ExecuteIn string `json:"ExecuteIn"` // CVM
GroupName string `json:"GroupName"` // test1
HookType string `json:"HookType"` // DEFAULT
ID int64 `json:"Id"` // 23
JenkinsFileFromType string `json:"JenkinsFileFromType"` // STATIC
JenkinsFilePath string `json:"JenkinsFilePath"` // Jenkinsfile
JenkinsFileStaticContent string `json:"JenkinsFileStaticContent"` // pipeline {
    agent any
    stages  {
        
        stage("检出") {
            steps {
                checkout(
                    [$class: 'GitSCM', branches: [[name: env.GIT_BUILD_REF]], 
                    userRemoteConfigs: [[url: env.GIT_REPO_URL, credentialsId: env.CREDENTIALS_ID]]]
                )
            }
        }

        stage("构建") {
            steps {
                echo "构建中..."
                sh 'docker version'
                // 请在这里放置您项目代码的单元测试调用过程，例如:
                // sh 'mvn package' // mvn 示例
                // sh 'make' // make 示例
                echo "构建完成."
              
                // 演示怎样产生构建物
                script{
                    def exists = fileExists 'README.md'
                    if (!exists) {
                        writeFile(file: 'README.md', text: 'Helloworld')
                    }
                }
                archiveArtifacts artifacts: 'README.md', fingerprint: true
              
                // archiveArtifacts artifacts: '**/target/*.jar', fingerprint: true // 收集构建产物
            }
        }

        stage("测试") {
            steps {
                echo "单元测试中..."
                // 请在这里放置您项目代码的单元测试调用过程，例如:
                // sh 'mvn test' // mvn 示例
                // sh 'make test' // make 示例
                echo "单元测试完成."
              
                // 演示怎么样生成测试报告
                writeFile(file: 'TEST-demo.junit4.AppTest.xml', text: '''
                    <testsuite name="demo.junit4.AppTest" time="0.053" tests="3" errors="0" skipped="0" failures="0">
                        <properties>
                        </properties>
                        <testcase name="testApp" classname="demo.junit4.AppTest" time="0.003"/>
                        <testcase name="test1" classname="demo.junit4.AppTest" time="0"/>
                        <testcase name="test2" classname="demo.junit4.AppTest" time="0"/>
                    </testsuite>
                ''')
                junit '*.xml'
              
                // junit 'target/surefire-reports/*.xml' // 收集单元测试报告的调用过程
            }
        }

        stage("部署") {
            steps {
                echo "部署中..."
                // 请在这里放置收集单元测试报告的调用过程，例如:
                // sh 'mvn tomcat7:deploy' // Maven tomcat7 插件示例：
                // sh './deploy.sh' // 自研部署脚本
                echo "部署完成"
            }
        }
    }
}

JobFromType string `json:"JobFromType"` // CODING
Name string `json:"Name"` // job-create
ProjectID int64 `json:"ProjectId"` // 450
ProjectName string `json:"ProjectName"` // test-1
ScheduleList []interface{} `json:"ScheduleList"` // <nil>
TriggerMethodList []string `json:"TriggerMethodList"` // MR_CHANGE
TriggerRemind string `json:"TriggerRemind"` // ALWAYS
UpdatedAt int64 `json:"UpdatedAt"` // 0
} `json:"Job"`
RequestID int64 `json:"RequestId,string"` // 1
} `json:"Response"`
}

// DescribeCodingCIJob 获取构建计划详情
func (c *Ci) DescribeCodingCIJob(req *DescribeCodingCIJobReq) error {
	return poster.Post(c.c, "DescribeCodingCIJob", req, nil)
}
