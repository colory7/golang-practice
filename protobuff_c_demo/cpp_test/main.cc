#include "msg.pb.h"
#include <sys/types.h>
#include <sys/socket.h>
#include <netinet/in.h>
#include <arpa/inet.h>
#include <stdio.h>
#include <cstring>
#include <iostream>
#include <string>

using namespace std;

int main(int argc, char *argv[]){

    int sk = socket(AF_INET, SOCK_STREAM, 0);

    struct sockaddr_in server;
    server.sin_family = AF_INET;
    server.sin_port = htons(2121); // 固定端口port
    server.sin_addr.s_addr = inet_addr("127.0.0.1"); // 固定ip

    connect(sk, (struct sockaddr*)&server, sizeof(server));

    msg::Msg sendMsg;
    sendMsg.set_msgid(0);
    sendMsg.set_msginfo("hello protobuf");
    sendMsg.set_msgfrom("alex");

    string sendData;
    sendMsg.SerializeToString(&sendData);
    int len = sendData.length();
    cout << "string len:" << len << endl;

    char *buff = new char[len + 1];
    memcpy(buff, sendData.c_str(), len);

    cout << "buff len:" << strlen(buff) << endl;
    //向服务段发送数据
    //在发送数据时一定要指明数据长度 防止中间有\0截断c风格字符串
    send(sk, buff, len, 0);
    close(sk);
    return 0;
}