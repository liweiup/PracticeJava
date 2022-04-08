package com.thread.test.thread;

import java.util.ArrayList;
import java.util.List;

public class ProducerConsumer {
    private static int MAX_SIZE = 3;
    private static final List<String> list = new ArrayList<>();

    public static void main(String[] args) {
        //创建生产者线程
        Producer producer = new Producer();
        //创建消费者线程
        Consumer consumer = new Consumer();
        //生产者线程开启
        producer.start();
        //消费者线程开启
//        consumer.start();
    }
    static class Producer extends Thread {
        @Override
        public void run() {
            while (true) {
                try {
                    //生产者 sleep 300ms, 消费者 sleep 500ms，模拟两者的处理能力不均衡
                    Thread.sleep(300);
                } catch (InterruptedException e1) {
                    e1.printStackTrace();
                }
                //第1步：获取队列对象的锁，与消费者持有的锁是同一把，保证线程安全
                synchronized (list) {
                    //第2步：判断缓冲区当前容量
                    //第2.1步：队列满了就不生产，等待
                    while (list.size() == MAX_SIZE) {
                        System.out.println("生产者 -> 缓冲区满了，等待消费...");
                        try {
                            //使用wait等待方法，内部会释放当前持有的锁
                            list.wait();
                            System.out.println("生产者 -> 休息了");
                        } catch (InterruptedException e) {
                            e.printStackTrace();
                        }
                    }
                    list.add("product");
                    System.out.println("生产者 -> 生产一个产品，当前队列大小：" + list.size());
                    //唤醒其他线程，这里其他线程就是指消费者线程
                    list.notify();
                }
            }
        }
    }
    static class Consumer extends Thread {
        @Override
        public void run() {
            //使用while循环执行run方法
            while (true) {
                try {
                    Thread.sleep(500);
                } catch (InterruptedException e1) {
                    e1.printStackTrace();
                }
                //第1步：获取队列对象的锁，与生产者持有的锁是同一把，保证线程安全
                synchronized (list) {
                    //第2步：判断缓冲区当前容量
                    //第2.1步：队列空了，等带
                    while (list.size() == 0) {
                        System.out.println("消费者 -> 缓冲区空了，等待生产...");
                        try {
                            //使用wait等待方法，内部会释放当前持有的锁
                            list.wait();
                        } catch (InterruptedException e) {
                            e.printStackTrace();
                        }
                    }
                    //第2.2步：队列不为空，消费一个产品
                    list.remove(0);
                    System.out.println("消费者 -> 消费一个产品，当前队列大小：" + list.size());

                    //唤醒其他线程，这里其他线程就是指生产者线程
                    list.notify();
                }
            }
        }
    }
}
