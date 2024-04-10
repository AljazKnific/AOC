import java.io.File;
import java.util.HashSet;
import java.util.LinkedList;
import java.util.Scanner;
import java.util.Set;

class Pair {
    int x;
    int y;

    public Pair(int x, int y) {
        this.x = x;
        this.y = y;
    }
}

class Element {
    int sum;
    Pair currPosition;
    // 0 - N 1 - E 2 - S 3 - W
    int steps;
    // max direction is three
    int dir;

    public Element(int x, int y, int sum, int steps, int dir) {
        this.sum = sum;
        this.currPosition = new Pair(x, y);
        this.steps = steps;
        this.dir = dir;
    }

}

public class Day {

    public static void addToQueue(LinkedList<Element> queue, int dir, int[][] seen, Element e, int[][] arr) {
        int x = e.currPosition.x;
        int y = e.currPosition.y;
        switch (dir) {
            // up one
            case 0:
                x--;
                break;
            // right one
            case 1:
                y++;
                break;
            case 2:
                x++;
                break;
            case 3:
                y--;
                break;
            default:
                System.err.println("Number needs to be [0,3]  " + dir);
                break;
        }

        // System.out.println("DIR: " + dir + " X: " + x + " Y: " + y);

        // we check if we are in the matrix
        if (x < 0 || x >= seen.length || y < 0 || y >= seen.length) {
            // System.out.println("FIRST OUT out of bounds");
            return;
        }

        if (seen[x][y] != 0 && seen[x][y] <= e.sum + arr[x][y]) {
            // System.out.println("SECOND OUT " + seen[x][y]);
            return;

        }

        if (e.dir == dir && e.steps == 3) {
            // System.out.println("THIRD OUT " + dir);
            return;
        }

        seen[x][y] = e.sum + arr[x][y];
        Element nE = new Element(x, y, e.sum + arr[x][y], 1, dir);

        if (nE.dir == e.dir) {
            nE.steps = e.dir + 1;
        }

        // postavi v pravi polozaj v vrsti, (gledamo glede na sum)

        int i = 0;
        for (i = 0; i < queue.size() && queue.get(i).sum < nE.sum; i++) {

        }
        // System.out.println(i);
        queue.add(i, nE);

    }

    public static void izpis(int[][] arr) {
        for (int i = 0; i < arr.length; i++) {
            for (int j = 0; j < arr.length; j++) {
                System.out.print(arr[i][j] + " ");
            }
            System.out.println();
        }
    }

    public static int algo(int[][] arr) {

        LinkedList<Element> queue = new LinkedList<>();
        int[][] seen = new int[arr.length][arr.length];
        Element e = new Element(0, 0, arr[0][0], 0, -1);
        queue.add(e);
        seen[0][0] = arr[0][0];

        while (!queue.isEmpty()) {
            Element first = queue.pop();
            // System.out.println(first.sum + " x: " + first.currPosition.x + " y: " +
            // first.currPosition.y);

            if (first.currPosition.x == arr.length - 1 && first.currPosition.y == arr.length - 1) {
                // izpis(seen);
                return first.sum;
            }

            // we can add for different elements -> 1 will always be discarded sometimes 2
            for (int i = 0; i < 4; i++) {
                addToQueue(queue, i, seen, first, arr);

            }

            // for (Element bs : queue) {
            // System.out.print(bs.sum + " x: " + bs.currPosition.x + " y: " +
            // bs.currPosition.y + " | ");
            // }
            // System.out.println();

        }

        return -42;
    }

    public static void main(String[] args) {
        try {
            int[][] arr = new int[13][13];
            // int[][] arr = new int[141][141];
            Scanner sc = new Scanner(new File("I.txt"));
            int y = 0;
            while (sc.hasNextLine()) {
                String[] data = sc.nextLine().split("");
                // System.out.println(data.length);
                for (int i = 0; i < data.length; i++) {
                    arr[y][i] = Integer.parseInt(data[i]);
                }
                y++;
            }
            System.out.println("Res: " + algo(arr));

            sc.close();
        } catch (Exception e) {
            System.err.println(e);
        }
    }
}