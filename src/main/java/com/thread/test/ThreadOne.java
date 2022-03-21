package com.thread.test;

import org.apache.commons.io.FileUtils;

import java.io.File;
import java.io.IOException;
import java.net.MalformedURLException;
import java.net.URL;

public class ThreadOne extends Thread{
    private String url;
    private String name;

    ThreadOne(String url,String name) {
        this.url = url;
        this.name = name;
    }

    public static void main(String[] args) {
        String url ="https://img2.baidu.com/it/u=3061412812,1793884610&fm=253&fmt=auto&app=138&f=PNG?w=500&h=312";
        Thread one = new ThreadOne(url,"/usr/local/www/java/src/xx1.jpg");
        Thread two = new ThreadOne(url,"/usr/local/www/java/src/xx2.jpg");
        Thread three = new ThreadOne(url,"/usr/local/www/java/src/xx3.jpg");
        Thread four = new ThreadOne(url,"/usr/local/www/java/src/xx4.jpg");
        one.start();
        two.start();
        three.start();
        four.start();
    }

    @Override
    public void run() {
        try {
            WebDownloader load = new WebDownloader();
            load.download(url,name);
            System.out.println("x下载了" + name);
        } catch (IOException e) {
            e.printStackTrace();
        }
    }

    static class WebDownloader {
        public void download(String url,String filePath) throws IOException {
            FileUtils.copyURLToFile(new URL(url),new File(filePath));
        }
    }
}