package com.thread.test;

public class ThreadClass extends  Thread{
    @Override
    public void run() {
        for (int i = 0; i < 20; i++) {
            System.out.println("我是自定义线程");
        }
    }

    public static void main(String[] args) {
        Thread thread_one = new ThreadClass();
        thread_one.start();
        for (int i = 0; i < 20; i++) {
            System.out.println("我是主线程");
        }
    }
}
