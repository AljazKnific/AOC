import java.io.File;
import java.util.Arrays;
import java.util.HashMap;
import java.util.HashSet;
import java.util.LinkedList;
import java.util.Scanner;

class Node {
    String val;
    Node left;
    Node right;

    public Node(String val) {
        this.val = val;
        this.left = null;
        this.right = null;
    }

    public Node addRight(Node right) {
        this.right = right;
        return right;
    }

    public Node addLeft(Node left) {
        this.left = left;
        return left;
    }
}

public class DayEight {

    public static Node valNode(HashSet<Node> map, String val, String left, String right) {
        Node l = null;
        Node r = null;
        Node main = null;

        for (Node node : map) {
            if (node.val.equals(val)) {
                main = node;
            }

            if (node.val.equals(left)) {
                l = node;
            }

            if (node.val.equals(right)) {
                r = node;
            }
        }

        if (l == null) {
            l = new Node(left);
        }

        if (r == null) {
            r = new Node(right);
        }

        if (main == null) {
            main = new Node(val);
        }

        if (left.equals(right)) {
            r = l;
        }

        // adding neighbours to the main node
        main.addLeft(l);
        main.addRight(r);

        // adding nodes to a map
        map.add(main);
        map.add(l);
        map.add(r);

        return main;

    }

    public static Long lcm(Long a, Long sum) {
        System.out.println("A: " + a + " B: " + sum);
        return ((a * sum) / gcd(a, sum));

    }

    public static Long gcd(Long a, Long b) {
        Long bigger = a;

        Long max = 1L;

        if (a < b) {
            bigger = b;
        } else if (a == b) {
            return a;
        }

        for (Long i = 1L; i <= bigger; i++) {
            if (a % i == 0 && b % i == 0) {
                max = i;
            }
        }

        return max;

    }

    public static void main(String[] args) {
        try {

            // Scanner sc = new Scanner(new File("Test.txt"));
            Scanner sc = new Scanner(new File("Input.txt"));

            String[] data = sc.nextLine().split("");
            sc.nextLine();

            HashSet<Node> map = new HashSet<>();
            LinkedList<Node> starting = new LinkedList<>();
            Node root = null;

            while (sc.hasNextLine()) {
                String[] arr = sc.nextLine().split(" = ");
                String left = arr[1].split(",")[0].substring(1);
                String right = arr[1].split(",")[1].substring(1, 4);
                // part one

                if (arr[0].equals("AAA")) {
                    root = valNode(map, arr[0], left, right);

                } else {
                    valNode(map, arr[0], left, right);
                }

                // part two
                /*
                 * if (arr[0].charAt(2) == 'A') {
                 * // System.out.println(arr[0]);
                 * starting.add(valNode(map, arr[0], left, right));
                 * } else {
                 * valNode(map, arr[0], left, right);
                 * }
                 */
            }

            Long sum = 0L;
            int index = 0;

            while (!root.val.equals("ZZZ")) {

                if (data[index].equals("R")) {
                    root = root.right;
                } else {
                    root = root.left;
                }

                index++;
                index %= data.length;
                sum++;
            }

            // Last lcm was evaluated with calculator, cause it wasw taking to long :)
            /*
             * Long sum2 = 1L;
             * 
             * for (Node node : starting) {
             * sum = 0L;
             * index = 0;
             * while (node.val.charAt(2) != 'Z') {
             * if (data[index].equals("R")) {
             * node = node.right;
             * } else {
             * node = node.left;
             * }
             * index++;
             * index %= data.length;
             * sum++;
             * 
             * }
             * 
             * sum2 = lcm(sum2, sum);
             * }
             * 
             * 
             */

            System.out.println("Sum: " + (sum - 1));
            // System.out.println("SUM2: " + sum2);

        } catch (Exception e) {
            System.err.println(e);
        }
    }

}
