package ci

import "github.com/juanjiTech/coding-open-api-sdk/openAPI/poster"

type TriggerCodingCIBuildReq struct {
	Action   string `json:"Action"`   // TriggerCodingCIBuild
	JobID    int64  `json:"JobId"`    // 23
	Revision string `json:"Revision"` // master
}

type TriggerCodingCIBuildResp struct {
Response struct {
Data struct {
Build struct {
Branch string `json:"Branch"` // 
Cause string `json:"Cause"` // coding 手动触发
CodingCIID string `json:"CodingCIId"` // cci-2-605077
CommitID string `json:"CommitId"` // 14c8e6e51ea01fc916dbcaf416f79a68f42c7634
CreatedAt int64 `json:"CreatedAt"` // 1.589084450366e+12
Duration int64 `json:"Duration"` // 0
FailedMessage string `json:"FailedMessage"` // 
ID int64 `json:"Id"` // 55
JenkinsFileContent string `json:"JenkinsFileContent"` // pipeline {
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

JobID int64 `json:"JobId"` // 23
Number int64 `json:"Number"` // 2
Status string `json:"Status"` // QUEUED
StatusNode string `json:"StatusNode"` // 
TestResult struct {
Duration int64 `json:"Duration"` // 0
Empty bool `json:"Empty"` // false
FailCount int64 `json:"FailCount"` // 0
PassCount int64 `json:"PassCount"` // 0
SkipCount int64 `json:"SkipCount"` // 0
TotalCount int64 `json:"TotalCount"` // 0
} `json:"TestResult"`
} `json:"Build"`
} `json:"Data"`
RequestID int64 `json:"RequestId,string"` // 1
} `json:"Response"`
}

// TriggerCodingCIBuild 触发构建
func (c *Ci) TriggerCodingCIBuild(req *TriggerCodingCIBuildReq) (resp TriggerCodingCIBuildResp, err error) {
	err = poster.Post(c.c, "TriggerCodingCIBuild", req, &resp)
	return
}
