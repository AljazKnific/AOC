import java.io.File;
import java.util.Arrays;
import java.util.HashSet;
import java.util.LinkedList;
import java.util.Scanner;

class Pair {
    int x;
    int y;

    public Pair(int x, int y) {
        this.x = x;
        this.y = y;
    }
}

class QuickerResult {
    int[] xCircle;
    int[] yCircle;
    LinkedList<Pair> set;
    char[][] matrix;

    public QuickerResult(int size) {
        this.xCircle = new int[size];
        this.yCircle = new int[size];
        this.set = new LinkedList<>();
        this.matrix = new char[size][size];
    }
}

class Day {

    public static char[][] fixPosN(char[][] a) {
        for (int i = 1; i < a.length; i++) {
            for (int j = 0; j < a.length; j++) {
                if (a[i][j] == 'O') {
                    int temp = i - 1;
                    while (temp > -1 && a[temp][j] == '.') {
                        a[temp + 1][j] = '.';
                        a[temp][j] = 'O';
                        temp--;
                    }
                }
            }
        }

        return a;
    }

    public static char[][] fixPosS(char[][] a) {
        for (int i = a.length - 2; i >= 0; i--) {
            for (int j = a.length - 1; j >= 0; j--) {
                if (a[i][j] == 'O') {
                    int temp = i + 1;
                    while (temp < a.length && a[temp][j] == '.') {
                        a[temp - 1][j] = '.';
                        a[temp][j] = 'O';
                        temp++;
                    }
                }
            }
        }

        return a;
    }

    public static char[][] fixPosE(char[][] a) {

        for (int i = a.length - 2; i >= 0; i--) {
            for (int j = a.length - 1; j >= 0; j--) {
                if (a[j][i] == 'O') {
                    int temp = i + 1;
                    while (temp < a.length && a[j][temp] == '.') {
                        a[j][temp - 1] = '.';
                        a[j][temp] = 'O';
                        temp++;
                    }
                }
            }
        }

        return a;
    }

    public static char[][] fixPosW(char[][] a) {
        for (int i = 1; i < a.length; i++) {
            for (int j = 0; j < a.length; j++) {
                if (a[j][i] == 'O') {
                    int temp = i - 1;
                    while (temp > -1 && a[j][temp] == '.') {
                        a[j][temp + 1] = '.';
                        a[j][temp] = 'O';
                        temp--;
                    }
                }
            }
        }

        return a;
    }

    public static QuickerResult getCoor(char[][] a) {
        QuickerResult q = new QuickerResult(a.length);
        for (int i = 0; i < a.length; i++) {
            for (int j = 0; j < a.length; j++) {
                q.matrix[i][j] = a[i][j];
                if (a[i][j] == 'O') {
                    q.xCircle[i]++;
                    q.yCircle[j]++;
                    q.set.add(new Pair(i, j));
                }
            }
        }

        return q;
    }

    public static void izpis(char[][] arr) {
        for (int i = 0; i < arr.length; i++) {
            for (int j = 0; j < arr.length; j++) {
                System.out.print(arr[i][j]);
            }
            System.out.println();
        }

    }

    public static int getSum(char[][] arr) {
        int sum = 0;
        int rock = 0;
        int[] list = new int[arr.length];

        for (int i = 0; i < arr.length; i++) {
            for (int j = 0; j < arr.length; j++) {
                switch (arr[i][j]) {
                    case 'O':
                        rock++;
                        sum += list[j];
                        list[j]++;
                        break;
                    case '.':
                        break;
                    case '#':
                        list[j] = i + 1;
                        break;
                    default:
                        System.err.println("WRONG INPUT");
                        break;
                }
            }
        }
        return (rock * arr.length - sum);
    }

    // This works -> also it's subroutines
    public static char[][] cycle(char[][] a) {

        return fixPosE(fixPosS(fixPosW(fixPosN(a))));
    }

    public static boolean areSame(QuickerResult r1, QuickerResult r2) {

        for (int i = 0; i < r1.xCircle.length; i++) {
            if (r1.xCircle[i] != r2.xCircle[i]) {
                return false;
            }
            if (r1.yCircle[i] != r2.yCircle[i]) {
                return false;
            }
        }

        for (int i = 0; i < r1.set.size(); i++) {
            Pair p1 = r1.set.get(i);
            Pair p2 = r2.set.get(i);
            if (p1.x != p2.x || p1.y != p2.y) {
                return false;
            }
        }

        return true;

    }

    public static int getSumPartTwo(int[] y) {
        int sum = 0;
        for (int i = 0; i < y.length; i++) {
            // System.out.println(y[i] + " * " + y.length + " - " + i);
            sum += (y[i] * (y.length - i));
        }
        // System.out.println();

        return sum;

    }

    public static void main(String[] args) {

        try {
            Scanner sc = new Scanner(new File("Input.txt"));
            char[][] list = new char[100][100];
            int pos = 0;
            while (sc.hasNextLine()) {
                String data = sc.nextLine();

                for (int i = 0; i < data.length(); i++) {
                    list[pos][i] = data.charAt(i);
                }

                pos++;
            }

            // Funkcije delujejo, problem nastane pri izracunu dela 2
            LinkedList<QuickerResult> grids = new LinkedList<>();
            QuickerResult tt = getCoor(list);
            grids.add(tt);

            int i = 0;
            int index = 0;
            for (i = 1; i < 1000000000; i++) {

                QuickerResult t1 = getCoor(cycle(grids.getLast().matrix));
                boolean t = false;
                // MAtrices are the same
                for (int j = 0; j < grids.size(); j++) {
                    if (areSame(grids.get(j), t1)) {
                        System.out.println(j);
                        index = j;
                        t = true;
                        break;

                    }
                }

                if (t) {
                    break;
                }

                grids.add(t1);

            }

            System.out
                    .println("SUM: " + getSumPartTwo(grids.get(((1000000000 - index) % (i - index)) + index).xCircle));

            sc.close();
        } catch (Exception e) {
            System.err.println(e);
        }
    }
}