package com.thread.test;

public class VolatileDemo {
    public static volatile boolean stop=true;

    public static void main(String[] args) throws InterruptedException {
        new Thread(()->{
            int i=0;
            while (stop){
                i++;
            }
            System.out.println("循环结束："+i);
        }).start();
        Thread.sleep(1000);
        stop=false;

        synchronized (new Object()) {

        }
    }
}

