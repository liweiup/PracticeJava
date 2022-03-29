package com.thread.test.thread;


import java.util.concurrent.*;

public class ThreadClass extends  Thread{
    public volatile int num = 1;
    @Override
    public void run() {
        for (int i = 0; i < 2000; i++) {
            System.out.println(Thread.currentThread().getName()+"我是自定义线程");
        }
    }

    public static void main(String[] args) {
        BlockingQueue<Runnable> workQueue = new LinkedBlockingQueue<>(20);
        RejectedExecutionHandler handler = new RejectedExecutionHandler() {
            @Override
            public void rejectedExecution(Runnable r, ThreadPoolExecutor executor) {
//                System.out.println(111);
            }
        };
        ThreadFactory factory = new ThreadFactory() {
            @Override
            public Thread newThread(Runnable r) {
                Thread t = new Thread(r, "my-thread-");
                return t;
            }
        };

        ThreadPoolExecutor thread = new ThreadPoolExecutor(1,3,2, TimeUnit.SECONDS,workQueue, factory,handler);
        Thread thread_one = new ThreadClass();
//        thread_one.start();
        thread.execute(thread_one);
        for (int i = 0; i < 2000; i++) {
            System.out.println("我是主线程");
        }
    }
}
