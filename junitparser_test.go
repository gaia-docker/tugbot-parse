package parse

import (
	log "github.com/Sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJavaJunitTestSuite(t *testing.T) {
	log.Info("TestJavaJunitTestSuite")
	suite, err := toJunitTestSuite([]byte(getMockJunit()))
	assert.NoError(t, err)
	assert.Equal(t, 5, len(suite.TestCases))
	assert.Equal(t, 5, len(suite.TestCases))
	assert.Equal(t, "5", suite.Total, "Total")
	assert.Equal(t, "1", suite.Failures, "Failures")
	assert.Equal(t, "0", suite.Errors, "Errors")
	assert.Equal(t, "0", suite.Skipped, "Skipped")
}

func TestMochaXunitTestSuite(t *testing.T) {
	log.Info("TestMochaXunitTestSuite")
	suite, err := toJunitTestSuite([]byte(getMockMochaXunit()))
	assert.NoError(t, err)
	assert.Equal(t, 5, len(suite.TestCases))
	assert.Equal(t, "5", suite.Total, "Total")
	assert.Equal(t, "1", suite.Failures, "Failures")
	assert.Equal(t, "1", suite.Errors, "Errors")
	assert.Equal(t, "0", suite.Skipped, "Skipped")
}

func TestMochaXunitTestSuiteWithSkippedTest(t *testing.T) {
	log.Info("TestMochaXunitTestSuiteWithSkippedTest")
	suite, err := toJunitTestSuite([]byte(getMockMochaXunitWithSkippedTest()))
	assert.NoError(t, err)
	assert.Equal(t, 5, len(suite.TestCases))
	assert.Equal(t, "5", suite.Total, "Total")
	assert.Equal(t, "4", suite.Failures, "Failures")
	assert.Equal(t, "4", suite.Errors, "Errors")
	assert.Equal(t, "1", suite.Skipped, "Skipped")
	log.Infof("%+v", suite)
}

func getMockJunit() string {

	// TEST-com.hp.mqm.testbox.service.TestResultServiceImplGherkinITCase.xml
	return `<?xml version="1.0" encoding="UTF-8"?>
	<testsuite name="com.hp.mqm.testbox.service.TestResultServiceImplGherkinITCase" time="39.202" tests="5" errors="0" skipped="0" failures="1">
  <properties>
    <property name="java.vendor" value="Oracle Corporation"/>
    <property name="TRACK" value="quick"/>
    <property name="sun.java.launcher" value="SUN_STANDARD"/>
    <property name="sun.management.compiler" value="HotSpot 64-Bit Tiered Compilers"/>
    <property name="os.name" value="Linux"/>
    <property name="sun.boot.class.path" value="/DATA/ads_slave/tools/hudson.model.JDK/OpenJDK-8u65/jre/lib/resources.jar:/DATA/ads_slave/tools/hudson.model.JDK/OpenJDK-8u65/jre/lib/rt.jar:/DATA/ads_slave/tools/hudson.model.JDK/OpenJDK-8u65/jre/lib/sunrsasign.jar:/DATA/ads_slave/tools/hudson.model.JDK/OpenJDK-8u65/jre/lib/jsse.jar:/DATA/ads_slave/tools/hudson.model.JDK/OpenJDK-8u65/jre/lib/jce.jar:/DATA/ads_slave/tools/hudson.model.JDK/OpenJDK-8u65/jre/lib/charsets.jar:/DATA/ads_slave/tools/hudson.model.JDK/OpenJDK-8u65/jre/lib/jfr.jar:/DATA/ads_slave/tools/hudson.model.JDK/OpenJDK-8u65/jre/classes"/>
    <property name="BUILD_NODE" value="mydtbld0138.hpeswlab.net"/>
    <property name="db.schema.name" value="mqm1472813603860"/>
    <property name="CUSTOM_JAVA" value="/DATA/ads_slave/tools/hudson.model.JDK/OpenJDK-8u65"/>
    <property name="java.vm.specification.vendor" value="Oracle Corporation"/>
    <property name="java.runtime.version" value="1.8.0_65-b17"/>
    <property name="archive_build_number" value="1477"/>
    <property name="SCM_BRANCH" value="master"/>
    <property name="user.name" value="ads_slave"/>
    <property name="mercury.td.sa_config_dir" value="target/generate-schema-out"/>
    <property name="guice.disable.misplaced.annotation.check" value="true"/>
    <property name="maven.repo.local" value="/DATA/ads_slave/workspace/MQM-System-Test-quick-master/.repository"/>
    <property name="user.language" value="en"/>
    <property name="sun.boot.library.path" value="/DATA/ads_slave/tools/hudson.model.JDK/OpenJDK-8u65/jre/lib/amd64"/>
    <property name="classworlds.conf" value="/DATA/ads_slave/tools/hudson.tasks.Maven_MavenInstallation/maven-3.0.5/bin/m2.conf"/>
    <property name="java.version" value="1.8.0_65"/>
    <property name="user.timezone" value="Asia/Jerusalem"/>
    <property name="sun.arch.data.model" value="64"/>
    <property name="jetty.port" value="${JETTY_PORT}"/>
    <property name="java.endorsed.dirs" value="/DATA/ads_slave/tools/hudson.model.JDK/OpenJDK-8u65/jre/lib/endorsed"/>
    <property name="sun.cpu.isalist" value=""/>
    <property name="sun.jnu.encoding" value="UTF-8"/>
    <property name="file.encoding.pkg" value="sun.io"/>
    <property name="MQM_VERSION" value="12.53.11"/>
    <property name="file.separator" value="/"/>
    <property name="java.specification.name" value="Java Platform API Specification"/>
    <property name="com.hp.mqm.infra.log.log4j.xml.location" value="/DATA/ads_slave/workspace/MQM-System-Test-quick-master/mqm/Server/mqm-integration-tests/target/classes/log4j.xml"/>
    <property name="java.class.version" value="52.0"/>
    <property name="user.country" value="US"/>
    <property name="securerandom.source" value="file:/dev/./urandom"/>
    <property name="ARCHIVE_JOB_NAME" value="MQM-job-Archive-quick-master"/>
    <property name="alm.home" value="/DATA/ads_slave/workspace/MQM-System-Test-quick-master/tmp/dev-home"/>
    <property name="java.home" value="/DATA/ads_slave/tools/hudson.model.JDK/OpenJDK-8u65/jre"/>
    <property name="jetty.stop.port" value="${JETTY_STOP_PORT}"/>
    <property name="java.vm.info" value="mixed mode"/>
    <property name="os.version" value="3.10.0-327.el7.x86_64"/>
    <property name="db.properties.file.path" value="/DATA/ads_slave/workspace/MQM-System-Test-quick-master/db/DBSchemasBuilder.ORACLE.properties"/>
    <property name="path.separator" value=":"/>
    <property name="java.vm.version" value="25.65-b01"/>
    <property name="UNIT_TEST" value="true"/>
    <property name="java.awt.printerjob" value="sun.print.PSPrinterJob"/>
    <property name="sun.io.unicode.encoding" value="UnicodeLittle"/>
    <property name="awt.toolkit" value="sun.awt.X11.XToolkit"/>
    <property name="workspace" value="/DATA/ads_slave/workspace/MQM-System-Test-quick-master"/>
    <property name="PACKAGE_VERSION" value="12.53.11-SNAPSHOT"/>
    <property name="user.home" value="/DATA/ads_slave/workspace/MQM-System-Test-quick-master/usr"/>
    <property name="VALIDATE_JS" value="true"/>
    <property name="java.specification.vendor" value="Oracle Corporation"/>
    <property name="java.library.path" value="/opt/HPCSS/HPSignClient/lunasa/lib:/usr/java/packages/lib/amd64:/usr/lib64:/lib64:/lib:/usr/lib"/>
    <property name="java.vendor.url" value="http://java.oracle.com/"/>
    <property name="ROOT_BUILD_NUMBER" value="13672"/>
    <property name="java.vm.vendor" value="Oracle Corporation"/>
    <property name="maven.home" value="/DATA/ads_slave/tools/hudson.tasks.Maven_MavenInstallation/maven-3.0.5"/>
    <property name="java.runtime.name" value="OpenJDK Runtime Environment"/>
    <property name="sun.java.command" value="org.codehaus.plexus.classworlds.launcher.Launcher -f mqm/Server/mqm-integration-tests/pom.xml -s /tmp/settings8238122921082597807.xml -DVALIDATE_JS=true -DSCM_BRANCH=master -DINDI_PACKAGE_VERSION=12.53.11-SNAPSHOT -DGIT_COMMIT=5d69fb87bc0581aea90966aedbd564ec733e2c83 -DROOT_BUILD_NUMBER=13672 -DARCHIVE_JOB_NAME=MQM-job-Archive-quick-master -DTRACK=quick -DMQM_VERSION=12.53.11 -DCUSTOM_JAVA=/DATA/ads_slave/tools/hudson.model.JDK/OpenJDK-8u65 -Darchive_build_number=1477 -DDOC_VERSION=12.53.11-SNAPSHOT -DBUILD_NODE=mydtbld0138.hpeswlab.net -DSONAR=true -DPACKAGE_VERSION=12.53.11-SNAPSHOT -DUNIT_TEST=true -DINTEGRATION_TEST=true -DCUSTOM_WORKSPACE=/DATA/ads_slave/workspace/MQM-Root-quick-master -DRepositoryFolder=/DATA/ads_slave/workspace/MQM-System-Test-quick-master/tmp -Dalm.home=/DATA/ads_slave/workspace/MQM-System-Test-quick-master/tmp/dev-home -Dworkspace=/DATA/ads_slave/workspace/MQM-System-Test-quick-master -Ddb.properties.file.path=/DATA/ads_slave/workspace/MQM-System-Test-quick-master/db/DBSchemasBuilder.ORACLE.properties -Drestapi.baseurl=http://localhost:${JETTY_PORT}/qcbin -Duser.home=/DATA/ads_slave/workspace/MQM-System-Test-quick-master/usr -Dmercury.td.sa_config_dir=target/generate-schema-out -Duser.dir=/DATA/ads_slave/workspace/MQM-System-Test-quick-master/tmp/dev-home -Ddb.schema.name=mqm1472813603860 -Djetty.stop.port=${JETTY_STOP_PORT} -Dcom.hp.mqm.infra.log.log4j.xml.location=/DATA/ads_slave/workspace/MQM-System-Test-quick-master/mqm/Server/mqm-integration-tests/target/classes/log4j.xml -Djetty.port=${JETTY_PORT} -Djava.io.tmpdir=/DATA/ads_slave/workspace/MQM-System-Test-quick-master/tmp -Ddb.solo.home=/DATA/ads_slave/tools/com.cloudbees.jenkins.plugins.customtools.CustomTool/DBSolo5 -Dmaven.repo.local=/DATA/ads_slave/workspace/MQM-System-Test-quick-master/.repository --batch-mode -nsu --errors --activate-profiles createSASchema,egg-test,itest-coverage,it-quick clean integration-test"/>
    <property name="java.class.path" value="/DATA/ads_slave/tools/hudson.tasks.Maven_MavenInstallation/maven-3.0.5/boot/plexus-classworlds-2.4.jar"/>
    <property name="RepositoryFolder" value="/DATA/ads_slave/workspace/MQM-System-Test-quick-master/tmp"/>
    <property name="java.vm.specification.name" value="Java Virtual Machine Specification"/>
    <property name="CUSTOM_WORKSPACE" value="/DATA/ads_slave/workspace/MQM-Root-quick-master"/>
    <property name="java.vm.specification.version" value="1.8"/>
    <property name="sun.cpu.endian" value="little"/>
    <property name="sun.os.patch.level" value="unknown"/>
    <property name="java.io.tmpdir" value="/DATA/ads_slave/workspace/MQM-System-Test-quick-master/tmp"/>
    <property name="java.vendor.url.bug" value="http://bugreport.sun.com/bugreport/"/>
    <property name="restapi.baseurl" value="http://localhost:${JETTY_PORT}/qcbin"/>
    <property name="java.awt.graphicsenv" value="sun.awt.X11GraphicsEnvironment"/>
    <property name="os.arch" value="amd64"/>
    <property name="INTEGRATION_TEST" value="true"/>
    <property name="java.ext.dirs" value="/DATA/ads_slave/tools/hudson.model.JDK/OpenJDK-8u65/jre/lib/ext:/usr/java/packages/lib/ext"/>
    <property name="INDI_PACKAGE_VERSION" value="12.53.11-SNAPSHOT"/>
    <property name="user.dir" value="/DATA/ads_slave/workspace/MQM-System-Test-quick-master/tmp/dev-home"/>
    <property name="line.separator" value="&#10;"/>
    <property name="java.vm.name" value="OpenJDK 64-Bit Server VM"/>
    <property name="GIT_COMMIT" value="5d69fb87bc0581aea90966aedbd564ec733e2c83"/>
    <property name="file.encoding" value="UTF-8"/>
    <property name="DOC_VERSION" value="12.53.11-SNAPSHOT"/>
    <property name="java.specification.version" value="1.8"/>
    <property name="SONAR" value="true"/>
    <property name="db.solo.home" value="/DATA/ads_slave/tools/com.cloudbees.jenkins.plugins.customtools.CustomTool/DBSolo5"/>
  </properties>
  <testcase name="testPersistTestResult_Gherkin_withoutTag_newTests" classname="com.hp.mqm.testbox.service.TestResultServiceImplGherkinITCase" time="9.584"/>
  <testcase name="testPersistTestResult_Gherkin_withTag_testExists" classname="com.hp.mqm.testbox.service.TestResultServiceImplGherkinITCase" time="12.198">
    <failure message="Expects to have one run for test expected:&lt;1&gt; but was:&lt;2&gt;" type="java.lang.AssertionError"><![CDATA[java.lang.AssertionError: Expects to have one run for test expected:<1> but was:<2>
	at org.junit.Assert.fail(Assert.java:88)
	at org.junit.Assert.failNotEquals(Assert.java:743)
	at org.junit.Assert.assertEquals(Assert.java:118)
	at org.junit.Assert.assertEquals(Assert.java:555)
	at com.hp.mqm.testbox.service.TestResultServiceImplGherkinITCase.assertAllGherkinTestEntities(TestResultServiceImplGherkinITCase.java:495)
	at com.hp.mqm.testbox.service.TestResultServiceImplGherkinITCase.testPersistTestResult_Gherkin_withTag_testExists(TestResultServiceImplGherkinITCase.java:339)
	at sun.reflect.NativeMethodAccessorImpl.invoke0(Native Method)
	at sun.reflect.NativeMethodAccessorImpl.invoke(NativeMethodAccessorImpl.java:62)
	at sun.reflect.DelegatingMethodAccessorImpl.invoke(DelegatingMethodAccessorImpl.java:43)
	at java.lang.reflect.Method.invoke(Method.java:497)
	at org.junit.runners.model.FrameworkMethod$1.runReflectiveCall(FrameworkMethod.java:47)
	at org.junit.internal.runners.model.ReflectiveCallable.run(ReflectiveCallable.java:12)
	at org.junit.runners.model.FrameworkMethod.invokeExplosively(FrameworkMethod.java:44)
	at org.junit.internal.runners.statements.InvokeMethod.evaluate(InvokeMethod.java:17)
	at org.junit.internal.runners.statements.RunBefores.evaluate(RunBefores.java:26)
	at org.springframework.test.context.junit4.statements.RunBeforeTestMethodCallbacks.evaluate(RunBeforeTestMethodCallbacks.java:75)
	at org.springframework.test.context.junit4.statements.RunAfterTestMethodCallbacks.evaluate(RunAfterTestMethodCallbacks.java:86)
	at org.springframework.test.context.junit4.statements.RunPrepareTestInstanceCallbacks.evaluate(RunPrepareTestInstanceCallbacks.java:64)
	at org.springframework.test.context.junit4.statements.SpringRepeat.evaluate(SpringRepeat.java:84)
	at org.springframework.test.context.junit4.statements.SpringFailOnTimeout.evaluate(SpringFailOnTimeout.java:87)
	at org.springframework.test.context.junit4.statements.ProfileValueChecker.evaluate(ProfileValueChecker.java:103)
	at org.junit.rules.TestWatcher$1.evaluate(TestWatcher.java:55)
	at org.junit.rules.RunRules.evaluate(RunRules.java:20)
	at org.junit.runners.ParentRunner.runLeaf(ParentRunner.java:271)
	at org.junit.runners.BlockJUnit4ClassRunner.runChild(BlockJUnit4ClassRunner.java:70)
	at junitparams.JUnitParamsRunner.runChild(JUnitParamsRunner.java:416)
	at junitparams.JUnitParamsRunner.runChild(JUnitParamsRunner.java:385)
	at org.junit.runners.ParentRunner$3.run(ParentRunner.java:238)
	at org.junit.runners.ParentRunner$1.schedule(ParentRunner.java:63)
	at org.junit.runners.ParentRunner.runChildren(ParentRunner.java:236)
	at org.junit.runners.ParentRunner.access$000(ParentRunner.java:53)
	at org.junit.runners.ParentRunner$2.evaluate(ParentRunner.java:229)
	at org.junit.internal.runners.statements.RunBefores.evaluate(RunBefores.java:26)
	at org.springframework.test.context.junit4.statements.RunBeforeTestClassCallbacks.evaluate(RunBeforeTestClassCallbacks.java:61)
	at org.springframework.test.context.junit4.statements.RunAfterTestClassCallbacks.evaluate(RunAfterTestClassCallbacks.java:70)
	at org.springframework.test.context.junit4.statements.ProfileValueChecker.evaluate(ProfileValueChecker.java:103)
	at org.springframework.test.context.junit4.rules.SpringClassRule$TestContextManagerCacheEvictor.evaluate(SpringClassRule.java:248)
	at org.junit.rules.RunRules.evaluate(RunRules.java:20)
	at org.junit.runners.ParentRunner.run(ParentRunner.java:309)
	at org.junit.runners.Suite.runChild(Suite.java:127)
	at org.junit.runners.Suite.runChild(Suite.java:26)
	at org.junit.runners.ParentRunner$3.run(ParentRunner.java:238)
	at org.junit.runners.ParentRunner$1.schedule(ParentRunner.java:63)
	at org.junit.runners.ParentRunner.runChildren(ParentRunner.java:236)
	at org.junit.runners.ParentRunner.access$000(ParentRunner.java:53)
	at org.junit.runners.ParentRunner$2.evaluate(ParentRunner.java:229)
	at org.junit.runners.ParentRunner.run(ParentRunner.java:309)
	at org.junit.runner.JUnitCore.run(JUnitCore.java:160)
	at org.junit.runner.JUnitCore.run(JUnitCore.java:138)
	at org.apache.maven.surefire.junitcore.JUnitCoreWrapper.createRequestAndRun(JUnitCoreWrapper.java:108)
	at org.apache.maven.surefire.junitcore.JUnitCoreWrapper.executeLazy(JUnitCoreWrapper.java:89)
	at org.apache.maven.surefire.junitcore.JUnitCoreWrapper.execute(JUnitCoreWrapper.java:58)
	at org.apache.maven.surefire.junitcore.JUnitCoreProvider.invoke(JUnitCoreProvider.java:144)
	at org.apache.maven.surefire.booter.ForkedBooter.invokeProviderInSameClassLoader(ForkedBooter.java:203)
	at org.apache.maven.surefire.booter.ForkedBooter.runSuitesInProcess(ForkedBooter.java:155)
	at org.apache.maven.surefire.booter.ForkedBooter.main(ForkedBooter.java:103)
]]></failure>
    <system-out><![CDATA[Setting user context to be:_SYSTEM_USER_
Setting user context to be:email4a697e8c-b40f-43fe-8712-4777e75d7cc1
Push duration: 491ms
Push duration: 391ms
Push duration: 268ms
]]></system-out>
  </testcase>
  <testcase name="testPersistTestResult_Gherkin_withoutTag_testsExist" classname="com.hp.mqm.testbox.service.TestResultServiceImplGherkinITCase" time="8.408"/>
  <testcase name="testPersistTestResult_Gherkin_withTag_testNotExists" classname="com.hp.mqm.testbox.service.TestResultServiceImplGherkinITCase" time="4.825"/>
  <testcase name="testPersistTestResult_Gherkin_with_outlineScenario" classname="com.hp.mqm.testbox.service.TestResultServiceImplGherkinITCase" time="4.187"/>
</testsuite>`
}

func getMockMochaXunit() string {

	return `<testsuite name="Mocha Tests" tests="5" failures="1" errors="1" skipped="0" timestamp="Mon, 25 Jul 2016 12:36:23 GMT" time="0.073">
			<testcase classname="Integration tests - voting" name="check database before voting" time="0.022"/>
			<testcase classname="Integration tests - voting" name="vote cats and verify" time="0.026"/>
			<testcase classname="Integration tests - voting" name="check database after voting" time="0">
				<failure>
					expected 0 to equal 1
					AssertionError: expected 0 to equal 1
					   at Assertion.assertEqual (node_modules/chai/lib/chai/core/assertions.js:487:12)
					   at Assertion.ctx.(anonymous function) [as equal] (node_modules/chai/lib/chai/utils/addMethod.js:41:25)
					   at null.&lt;anonymous&gt; (specs/e2e/voting-e2e-test.js:82:34)
					   at Query.handleReadyForQuery (node_modules/pg/lib/query.js:114:10)
					   at null.&lt;anonymous&gt; (node_modules/pg/lib/client.js:172:19)
					   at Socket.&lt;anonymous&gt; (node_modules/pg/lib/connection.js:121:12)
					   at readableAddChunk (_stream_readable.js:153:18)
					   at Socket.Readable.push (_stream_readable.js:111:10)
					   at TCP.onread (net.js:534:20)
				</failure>
			</testcase>
			<testcase classname="Functional tests - voting" name="open ui and check title" time="0.008"/>
			<testcase classname="Functional tests - voting" name="vote cats" time="0.005"/>
		</testsuite>`
}

func getMockMochaXunit2() string {

	return `<testsuite name="Mocha Tests" tests="5" failures="1" errors="1" skipped="0" timestamp="Mon, 25 Jul 2016 14:44:37 GMT" time="0.149">
		<testcase classname="Integration tests - voting" name="check database before voting" time="0.047"/>
		<testcase classname="Integration tests - voting" name="vote cats and verify" time="0.045"/>
		<testcase classname="Integration tests - voting" name="check database after voting" time="0.013"/>
		<testcase classname="Functional tests - voting" name="open ui and check title" time="0"><failure>expected -1 to be above -1
		AssertionError: expected -1 to be above -1
		   at Assertion.assertAbove (node_modules/chai/lib/chai/core/assertions.js:571:12)
		   at Assertion.ctx.(anonymous function) [as above] (node_modules/chai/lib/chai/utils/addMethod.js:41:25)
		   at Request._callback (specs/functional/vote-page-test.js:27:97)
		   at Request.self.callback (node_modules/request/request.js:187:22)
		   at Request.&lt;anonymous&gt; (node_modules/request/request.js:1044:10)
		   at IncomingMessage.&lt;anonymous&gt; (node_modules/request/request.js:965:12)
		   at endReadableNT (_stream_readable.js:913:12)
		   at _combinedTickCallback (internal/process/next_tick.js:74:11)
		   at process._tickCallback (internal/process/next_tick.js:98:9)</failure></testcase>
		<testcase classname="Functional tests - voting" name="vote cats" time="0.01"/>
		</testsuite>`
}

func getMockMochaXunitWithSkippedTest() string {

	return `
		<testsuite name="Mocha Tests" tests="5" failures="4" errors="4" skipped="1" timestamp="Wed, 27 Jul 2016 09:37:27 GMT" time="4.045">
			<testcase classname="Integration tests - voting" name="check database before voting" time="2.002">
				<failure>timeout of 2000ms exceeded. Ensure the done() callback is being called in this test.
					Error: timeout of 2000ms exceeded. Ensure the done() callback is being called in this test.</failure></testcase>
					<testcase classname="Integration tests - voting" name="vote cats and verify" time="0"><failure>Cannot read property 'statusCode' of undefined
					TypeError: Cannot read property 'statusCode' of undefined
					    at Request._callback (specs/e2e/voting-e2e-test.js:50:18)
					    at self.callback (node_modules/request/request.js:187:22)
					    at Request.onRequestError (node_modules/request/request.js:813:8)
					    at Socket.socketErrorListener (_http_client.js:306:9)
					    at emitErrorNT (net.js:1265:8)
					    at _combinedTickCallback (internal/process/next_tick.js:74:11)
					    at process._tickCallback (internal/process/next_tick.js:98:9)
			    	</failure>
			    </testcase>
			<testcase classname="Integration tests - voting" name="check database after voting" time="2.002"><failure>timeout of 2000ms exceeded. Ensure the done() callback is being called in this test.
			Error: timeout of 2000ms exceeded. Ensure the done() callback is being called in this test.</failure></testcase>
			<testcase classname="Functional tests - voting" name="open ui and check title" time="0"><failure>Cannot read property 'statusCode' of undefined
			TypeError: Cannot read property 'statusCode' of undefined
			    at Request._callback (specs/functional/vote-page-test.js:24:18)
			    at self.callback (node_modules/request/request.js:187:22)
			    at Request.onRequestError (node_modules/request/request.js:813:8)
			    at Socket.socketErrorListener (_http_client.js:306:9)
			    at emitErrorNT (net.js:1265:8)
			    at _combinedTickCallback (internal/process/next_tick.js:74:11)
			    at process._tickCallback (internal/process/next_tick.js:98:9)</failure></testcase>
			<testcase classname="Functional tests - voting" name="vote cats" time="0"><skipped>aaa</skipped></testcase>
		</testsuite>`
}
