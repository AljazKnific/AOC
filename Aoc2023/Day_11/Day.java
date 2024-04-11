import java.io.File;
import java.util.HashSet;
import java.util.LinkedList;
import java.util.Scanner;

import org.w3c.dom.css.Counter;

class Pair {
    int x;
    int y;

    public Pair(int x, int y) {
        this.x = x;
        this.y = y;
    }

    public Pair(int x, int y, int xM, int yM, int multi) {
        // System.out.println("X: " + x + " Y: " + y + " xM: " + xM + " yM: " + yM);
        // System.out.println((x - xM) + xM * multi);
        // System.out.println((y - yM) + yM * multi);

        this.x = (x - xM) + xM * multi;
        this.y = (y - yM) + yM * multi;
    }
}

class Day {

    public static Long sumDistance(HashSet<Pair> set) {
        HashSet<Pair> checked = new HashSet<>();
        Long sum = 0L;

        while (checked.size() != set.size()) {
            for (Pair pair : set) {
                // System.out.println("PAIR: " + pair.x + " : " + pair.y);
                for (Pair pair2 : set) {
                    // System.out.println(pair2.x + " : " + pair2.y);
                    if (!checked.contains(pair2) && !pair.equals(pair2)) {
                        sum += (Math.abs(pair.x - pair2.x) + Math.abs(pair.y - pair2.y));
                    }

                }
                checked.add(pair);
            }
        }

        return sum;
    }

    public static int numsOfI(HashSet<Integer> set, int x) {
        int sum = 0;

        for (Integer integer : set) {
            if (x > integer) {
                sum++;
                // System.out.println("X: " + x + " > " + integer);
            }
        }
        // System.out.println("Iscem: " + x + " st: " + sum);

        return sum;

    }

    public static void main(String[] args) {
        try {
            File inp = new File("Input.txt");
            Scanner sc = new Scanner(inp);

            LinkedList<LinkedList<Character>> list = new LinkedList<>();
            HashSet<Pair> set = new HashSet<>();
            HashSet<Integer> xCounter = new HashSet<>();
            HashSet<Integer> yCounter = new HashSet<>();
            // change based on the input
            int size = 140;
            int[] counter = new int[size];

            int index = 0;
            boolean space = true;

            while (sc.hasNextLine()) {
                LinkedList<Character> l = new LinkedList<>();
                list.add(l);
                String data = sc.nextLine();
                space = true;
                for (int i = 0; i < data.length(); i++) {
                    char c = data.charAt(i);
                    space &= (c == '.');
                    if (c == '.') {
                        counter[i]++;
                    }
                    l.add(c);
                }

                if (space) {
                    // adding x- indices of empty rows
                    xCounter.add(index);
                }

                index++;
            }

            for (int i = 0; i < counter.length; i++) {
                if (counter[i] == size) {
                    // adding y-indicies of empty columns
                    // System.out.println(i);
                    yCounter.add(i);

                }
            }

            /*
             * 
             * System.out.println("XCounter");
             * for (Integer i : xCounter) {
             * System.out.print(i + " ");
             * }
             * System.out.println();
             * System.out.println("yCounter");
             * for (Integer i : yCounter) {
             * System.out.print(i + " ");
             * }
             */

            for (int i = 0; i < list.size(); i++) {
                LinkedList<Character> l = list.get(i);
                for (int j = 0; j < l.size(); j++) {
                    if (l.get(j) == '#') {
                        set.add(new Pair(i, j, numsOfI(xCounter, i), numsOfI(yCounter, j), 1000000));
                    }
                }
            }

            System.out.println(sumDistance(set));

            /*
             * for (LinkedList<Character> linkedList : list) {
             * for (Character c : linkedList) {
             * System.out.print(c);
             * }
             * System.out.println();
             * }
             */
            sc.close();
        } catch (Exception e) {
            // TODO: handle exception
            System.err.println(e);
        }
    }
}