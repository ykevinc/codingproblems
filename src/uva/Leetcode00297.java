import java.io.*;
import java.util.*;

class Leetcode00297 {
  static class Node {
    private Node l;
    private Node r;
    private int data;
    public Node(int d) {data = d;}
  }
  
  static class Tree {
    private Node root;
    
    public Node insert(int d) {
      root = insert(root, d);
      return root;
    }
    private Node insert(Node node, int d) {
      if (node == null) {
        return new Node(d);
      } else if (node.data > d) {
        node.l = insert(node.l, d);
      } else if (node.data < d) {
        node.r = insert(node.r, d);
      } else {
        //ignore
      }
      return node;
    }
    public String serialize() {
      StringBuilder sb = new StringBuilder();
      serialize(root, sb);
      return sb.toString();
    }
    private void serialize(Node n, StringBuilder sb) {
      if (n == null) {
        sb.append("null");
        return;
      }
      sb.append(n.data);
      sb.append(" ");
      int c = 0;
      if (n.l != null) {
        c++;
      }
      if (n.r != null) {
        c++;
      }
      sb.append(c);
      sb.append(" ");
      if (n.l != null) {
        serialize(n.l, sb);
      }
      if (n.r != null) {
        serialize(n.r, sb);
      }
    }
    public static Tree deserialize(String s) {
      Scanner sr = new Scanner(s);
      Node newNode = deserialize(sr);
      Tree t = new Tree();
      t.root = newNode;
      return t;
    }
    public static Node deserialize(Scanner sr) {
      if (sr.hasNext()) {
        int data = sr.nextInt();
        Node node = new Node(data);
        int children = sr.nextInt();
        if (children == 2) {
          node.l = deserialize(sr);
          node.r = deserialize(sr);
        } else if (children == 1) {
          Node one = deserialize(sr);
          if(data > one.data) {
            node.l = one;
          } else if (data < one.data) {
            node.r = one;
          }
        }
        return node;
      }
      return null;
    }
  }
  
  public static void main(String[] args) {
    Tree t = new Tree();
    t.insert(3);
    t.insert(5);
    System.out.println(t.serialize());
    System.out.println(Tree.deserialize(t.serialize()).serialize());
  }
}
