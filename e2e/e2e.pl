#!/usr/bin/perl

$host = "localhost";
$port = 8080;

$status = `curl -s -o /dev/null -w "%{http_code}" http://$host:$port?userid=1`;
$status == 200 || die("userid=1 failed");
$status = `curl -s -o /dev/null -w "%{http_code}" http://$host:$port?userid=2`;
$status == 404 || die("userid=2 failed");
$status = `curl -s -o /dev/null -w "%{http_code}" http://$host:$port?userid=foo`;
$status == 400 || die("userid=foo failed");
$status = `curl -s -o /dev/null -w "%{http_code}" http://$host:$port`;
$status == 400 || die("no userid failed");
$json =  `curl -s http://$host:$port?userid=1`;
($json eq "[1]\n") == 1 || die("json failed");

printf "Well Done!\n";
