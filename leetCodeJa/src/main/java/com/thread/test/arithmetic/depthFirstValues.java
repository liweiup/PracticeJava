package com.thread.test.arithmetic;

import java.sql.Array;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.Stack;

class Node{
    public String val;
    public Node left;
    public Node right;

    public Node(String val) {
        this.val = val;
        this.left = null;
        this.right = null;
    }
    @Override
    public String toString() {
        return "Node{" +
                "val='" + val + '\'' +
                ", left=" + left +
                ", right=" + right +
                '}';
    }
}

public class depthFirstValues {

    public static void main(String[] args) {
        Node a = new Node("a");
        Node b = new Node("b");
        Node c = new Node("c");
        Node d = new Node("d");
        Node e = new Node("e");
        Node f = new Node("f");
        a.left = b;
        a.right = c;
        b.left = d;
        b.right = e;
        c.right = f;
        Stack<Node> s = new Stack<>();
        ArrayList<String> sites = new ArrayList<>();
        s.push(a);
//        while (s.size() > 0) {
//            Node current = s.pop();
//            sites.add(current.val);
//            if (current.right != null) {
//                s.push(current.right);
//            }
//            if (current.left != null) {
//                s.push(current.left);
//            }
//        }
//        System.out.println(sites);
        digui(a);
    }
    public static void digui(Node root) {
        if (root == null) {
            return;
        }
        System.out.println(root.val);
        digui(root.left);
        digui(root.right);
    }
}
