package com.thread.test;

import java.util.HashMap;
import java.util.concurrent.ConcurrentHashMap;
import java.util.concurrent.locks.ReentrantLock;

public class ConCurreentHashMap001 implements Runnable{
    private ConcurrentHashMap<String,String> c1;
//    private HashMap<String,String> c1;

    public ConCurreentHashMap001(ConcurrentHashMap<String, String> c1) {
        this.c1 = c1;
    }

    @Override
    public void run() {
        ReentrantLock lock = new ReentrantLock();
        for (int i = 0; i < 20; i++) {
            if (Thread.currentThread().getName().equals("wowo")) {
                System.out.println(Thread.currentThread().getName()+"============push"+c1.toString() + c1.put("11","111"));
            }
            if (Thread.currentThread().getName().equals("wowo1")) {
                System.out.println(Thread.currentThread().getName()+"============push"+c1.toString() + c1.put("11","2232"));
            }
            try {
                Thread.sleep(1000);
            } catch (InterruptedException e) {
                e.printStackTrace();
            }

        }
    }

    public static void main(String[] args) {
        ConcurrentHashMap<String,String> c1 = new ConcurrentHashMap<String,String>();
        c1.put("11","33322");
        ConCurreentHashMap001 conCurreentHashMap001 = new ConCurreentHashMap001(c1);
        new Thread(conCurreentHashMap001,"wowo").start();
        new Thread(conCurreentHashMap001,"wowo1").start();
//        System.out.println(conCurreentHashMap001.c1);
    }
}
