import java.io.File;
import java.util.HashSet;
import java.util.Scanner;

class Pair {
    int x;
    int y;

    public Pair(int x, int y) {
        this.x = x;
        this.y = y;
    }
}

public class Day {

    public static void izpis(char[][] arr) {
        for (int i = 0; i < arr.length; i++) {
            for (int j = 0; j < arr.length; j++) {

                System.out.print(arr[i][j]);
            }
            System.out.println();
        }
    }

    public static boolean isIn(Pair p, HashSet<Pair> set) {
        for (Pair pair : set) {
            if (p.x == pair.x && p.y == pair.y) {
                return true;
            }
        }
        return false;
    }

    public static void addNew(Pair p, char[][] arr, HashSet<Pair> newSet) {
        // north
        if (p.x > 0 && arr[p.x - 1][p.y] != '#') {
            Pair temp = new Pair(p.x - 1, p.y);
            if (!isIn(temp, newSet)) {
                newSet.add(temp);
            }
        }

        // south
        if (p.x < arr.length - 1 && arr[p.x + 1][p.y] != '#') {
            Pair temp = new Pair(p.x + 1, p.y);
            if (!isIn(temp, newSet)) {
                newSet.add(temp);
            }
        }

        // west
        if (p.y > 0 && arr[p.x][p.y - 1] != '#') {
            Pair temp = new Pair(p.x, p.y - 1);
            if (!isIn(temp, newSet)) {
                newSet.add(temp);
            }
        }

        // east
        if (p.y < arr.length - 1 && arr[p.x][p.y + 1] != '#') {
            Pair temp = new Pair(p.x, p.y + 1);
            if (!isIn(temp, newSet)) {
                newSet.add(temp);
            }
        }
    }

    public static char[][] createNewChar(char[][] arr) {
        char[][] t = new char[arr.length][arr.length];

        for (int i = 0; i < t.length; i++) {
            for (int j = 0; j < t.length; j++) {
                t[i][j] = arr[i][j];
            }
        }

        return t;
    }

    public static void izpis2(char[][] arr, HashSet<Pair> x) {
        for (Pair pair : x) {
            arr[pair.x][pair.y] = 'O';
        }

        izpis(arr);
    }

    public static void main(String[] args) {
        try {
            Scanner sc = new Scanner(new File("I.txt"));
            char[][] arr = new char[131][131];
            int index = 0;
            HashSet<Pair> set = new HashSet<>();
            Pair start = null;
            while (sc.hasNextLine()) {
                String data = sc.nextLine();
                for (int i = 0; i < data.length(); i++) {
                    arr[index][i] = data.charAt(i);
                    if (arr[index][i] == 'S') {
                        start = new Pair(index, i);
                    }
                }

                index++;
            }

            // izpis(arr);
            HashSet<Pair> newsSet = new HashSet<>();
            set.add(start);
            int a = 0;
            for (int i = 0; i < 64; i++) {
                a = 0;
                for (Pair pair : set) {
                    addNew(pair, arr, newsSet);
                }

                set = newsSet;
                if (!isIn(start, newsSet)) {
                    a = 1;
                }

                System.out.println("SIZE: " + (set.size() + a));
                // char[][] temp = createNewChar(arr);
                // izpis2(temp, newsSet);
                newsSet = new HashSet<>();
            }

            System.out.println("SIZE: " + (set.size() + a));

            sc.close();

        } catch (Exception e) {
            System.err.println(e);
        }
    }
}
