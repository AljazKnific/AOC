import java.io.File;
import java.util.Arrays;
import java.util.HashSet;
import java.util.Scanner;

public class Day {
    public static class Group {
        char c;
        HashSet<FROM> set;

        public Group(char c) {
            this.c = c;
            this.set = new HashSet<>();
        }

        public void reset() {
            this.set.clear();
        }
    }

    enum FROM {
        SOUTH,
        EAST,
        WEST,
        NORTH
    }

    public static int izpis(boolean[][] arr, Group[][] b) {
        int sum = 0;
        for (int i = 0; i < arr.length; i++) {
            for (int j = 0; j < arr.length; j++) {
                if (arr[i][j]) {
                    sum++;
                }
                b[i][j].reset();
                arr[i][j] = false;
            }
        }

        return sum;

    }

    public static boolean checkSet(FROM l, HashSet<FROM> set) {
        if (set.size() == 4) {
            return true;
        }

        for (FROM from : set) {
            if (from == l) {
                return true;
            }
        }

        return false;
    }

    public static void addAll(HashSet<FROM> set) {
        set.add(FROM.NORTH);
        set.add(FROM.EAST);
        set.add(FROM.WEST);
        set.add(FROM.SOUTH);
    }

    public static void addTwo(HashSet<FROM> set, FROM v1, FROM v2) {
        set.add(v2);
        set.add(v1);
    }

    public static void rec(int x, int y, FROM loc, Group[][] arr, boolean[][] check) {
        // ustavitveni pogoj
        if (x < 0 || x >= arr.length) {
            return;
        }
        if (y < 0 || y >= arr.length) {
            return;
        }
        // fix this -> / in \ sta lahko zadeta iz razlicnih smeri in proizvedeta druge
        // rezultate
        if (arr[x][y].c == '|' && arr[x][y].c == '-' && check[x][y]
                && arr[x][y].set.size() == 4) {
            return;
        }

        if (arr[x][y].c == '/' && check[x][y] && checkSet(loc, arr[x][y].set)) {
            return;
        }

        if (arr[x][y].c == '\\' && check[x][y] && checkSet(loc, arr[x][y].set)) {
            return;
        }

        // preglej katere smeri so v setu v primeru \ in /

        check[x][y] = true;

        switch (loc) {
            case NORTH:
                switch (arr[x][y].c) {
                    case '-':
                        addAll(arr[x][y].set);
                        rec(x, y + 1, FROM.WEST, arr, check);
                        rec(x, y - 1, FROM.EAST, arr, check);
                        break;
                    case '/':
                        addTwo(arr[x][y].set, FROM.NORTH, FROM.WEST);
                        rec(x, y - 1, FROM.EAST, arr, check);
                        break;
                    case '\\':
                        addTwo(arr[x][y].set, FROM.NORTH, FROM.EAST);
                        rec(x, y + 1, FROM.WEST, arr, check);
                        break;
                    case '|':
                        addAll(arr[x][y].set);
                        rec(x + 1, y, loc, arr, check);
                        break;
                    default:
                        // gremo samo eno navzdol
                        rec(x + 1, y, loc, arr, check);
                        break;
                }
                break;

            case SOUTH:
                switch (arr[x][y].c) {
                    case '-':
                        addAll(arr[x][y].set);
                        rec(x, y + 1, FROM.WEST, arr, check);
                        rec(x, y - 1, FROM.EAST, arr, check);
                        break;
                    case '/':
                        addTwo(arr[x][y].set, FROM.SOUTH, FROM.EAST);
                        rec(x, y + 1, FROM.WEST, arr, check);
                        break;
                    case '\\':
                        addTwo(arr[x][y].set, FROM.SOUTH, FROM.WEST);
                        rec(x, y - 1, FROM.EAST, arr, check);
                        break;
                    case '|':
                        addAll(arr[x][y].set);
                        rec(x - 1, y, loc, arr, check);
                        break;
                    default:
                        rec(x - 1, y, loc, arr, check);
                        break;
                }
                break;
            case WEST:
                switch (arr[x][y].c) {
                    case '|':
                        addAll(arr[x][y].set);
                        rec(x - 1, y, FROM.SOUTH, arr, check);
                        rec(x + 1, y, FROM.NORTH, arr, check);
                        break;
                    case '/':
                        addTwo(arr[x][y].set, FROM.NORTH, FROM.WEST);
                        rec(x - 1, y, FROM.SOUTH, arr, check);
                        break;
                    case '\\':
                        addTwo(arr[x][y].set, FROM.SOUTH, FROM.WEST);
                        rec(x + 1, y, FROM.NORTH, arr, check);
                        break;
                    case '-':
                        addAll(arr[x][y].set);
                        rec(x, y + 1, loc, arr, check);
                        break;
                    default:
                        rec(x, y + 1, loc, arr, check);
                        break;
                }
                break;
            case EAST:
                switch (arr[x][y].c) {
                    case '|':
                        addAll(arr[x][y].set);
                        rec(x - 1, y, FROM.SOUTH, arr, check);
                        rec(x + 1, y, FROM.NORTH, arr, check);
                        break;
                    case '/':
                        addTwo(arr[x][y].set, FROM.SOUTH, FROM.EAST);
                        rec(x + 1, y, FROM.NORTH, arr, check);
                        break;
                    case '\\':
                        addTwo(arr[x][y].set, FROM.NORTH, FROM.EAST);
                        rec(x - 1, y, FROM.SOUTH, arr, check);
                        break;
                    case '-':
                        addAll(arr[x][y].set);
                        rec(x, y - 1, loc, arr, check);
                        break;
                    default:
                        rec(x, y - 1, loc, arr, check);
                        break;
                }
                break;
            default:
                break;
        }
    }

    public static void main(String[] args) {

        try {
            Scanner sc = new Scanner(new File("I.txt"));
            boolean[][] check = new boolean[110][110];
            Group[][] array = new Group[110][110];

            int id = 0;

            while (sc.hasNextLine()) {
                String data = sc.nextLine();
                for (int i = 0; i < data.length(); i++) {
                    array[id][i] = new Group(data.charAt(i));
                }
                id++;
            }

            rec(0, 0, FROM.WEST, array, check);

            // izpis(array);

            System.out.println("SUM: " + izpis(check, array));
            int max = 0;
            // zgornja vrstica
            System.out.println("CHECK TOP ROW");
            for (int i = 0; i < array.length; i++) {
                rec(0, i, FROM.NORTH, array, check);
                int res = izpis(check, array);
                if (max < res) {
                    max = res;
                }
            }

            System.out.println("CHECK BOTTOM ROW");

            for (int i = 0; i < array.length; i++) {
                rec(array.length - 1, i, FROM.SOUTH, array, check);
                int res = izpis(check, array);
                if (max < res) {
                    max = res;
                }
            }

            System.out.println("CHECK LEFT COLOUMN");

            for (int i = 1; i < array.length - 1; i++) {
                rec(i, 0, FROM.WEST, array, check);
                int res = izpis(check, array);
                if (max < res) {
                    max = res;
                }
            }

            System.out.println("CHECK RIGHT COLOUMN");

            for (int i = 1; i < array.length - 1; i++) {
                rec(i, array.length - 1, FROM.EAST, array, check);
                int res = izpis(check, array);
                if (max < res) {
                    max = res;
                }
            }

            System.out.println("MAX: " + max);

            sc.close();
        } catch (Exception e) {
            System.err.println(e);
        }
    }

}
