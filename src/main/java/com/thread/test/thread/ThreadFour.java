package com.thread.test.thread;

import java.util.concurrent.atomic.AtomicInteger;

public class ThreadFour implements Runnable{
//    private int stockNum = 10;
    private AtomicInteger stockNum = new AtomicInteger(1);
    @Override
    public void run() {
        while (stockNum.intValue() <= 10) {
            System.out.println(Thread.currentThread().getName().toString() +"抢到了"+ stockNum.getAndIncrement());
        }
    }

    public static void main(String[] args) {
        ThreadFour threadFour = new ThreadFour();
        new Thread(threadFour,"xiaoli").start();
        new Thread(threadFour,"xiaozhang").start();
        new Thread(threadFour,"huangniu").start();
    }
}
