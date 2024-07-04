package day3

// - HTTP 웹서버(no DB) 부하 테스트 (with JMeter)
// https://jmeter.apache.org/download_jmeter.cgi

// wget https://downloads.apache.org//jmeter/binaries/apache-jmeter-5.6.3.tgz
// curl -O https://downloads.apache.org/jmeter/binaries/apache-jmeter-5.5.tgz

// tar -xvzf apache-jmeter-5.5.tgz

// sudo mv apache-jmeter-5.6.3 /usr/local/bin/jmeter

// echo "export PATH=\$PATH:/usr/local/bin/jmeter/bin" >> ~/.zshrc
// source ~/.zshrc

// jmeter -v
func Load() {

}

// https://github.com/apache/jmeter/issues/6083

// brew list --cask | grep corretto
// brew update
// brew install jmeter
// jmeter --version
// JAVA_HOME=/Library/Java/JavaVirtualMachines/amazon-corretto-22.jdk/Contents/Home/ exec /opt/homebrew/Cellar/jmeter/5.6.3/libexec/bin/jmeter
// alias rjm="JAVA_HOME=/Library/Java/JavaVirtualMachines/amazon-corretto-22.jdk/Contents/Home/ exec /opt/homebrew/Cellar/jmeter/5.6.3/libexec/bin/jmeter"

// https://velog.io/@zinna_1109/JMeter-Jmeter-Http-Request-%EC%9A%94%EC%B2%AD-%EB%B0%A9%EB%B2%95-%EC%84%B1%EB%8A%A5-%ED%85%8C%EC%8A%A4%ED%8A%B8
