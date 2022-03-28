package com.thread.test;

public class threadRunnable implements Runnable{

    public static void main(String[] args) {
        threadRunnable threadRunnable = new threadRunnable();
        new Thread(threadRunnable).start();
        for (int i = 0; i < 200; i++) {
            System.out.println("我是主线程");
        }
    }
    @Override
    public void run() {
        for (int i = 0; i < 200; i++) {
            System.out.println("我是自定义线程");
        }
    }
}
