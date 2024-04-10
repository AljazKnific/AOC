import java.io.File;
import java.util.Arrays;
import java.util.LinkedList;
import java.util.Scanner;

// idea: double queue where the first one checks if we were aleready there and  the second one for searching elements 

class Pair {
    int x;
    int y;

    public Pair(int x, int y) {
        this.x = x;
        this.y = y;
    }
}

class DayTeen {

    public static boolean checkUP(String s1, String s2) {
        char c1 = s1.charAt(0);
        char c2 = s2.charAt(0);

        boolean b1 = false;
        boolean b2 = false;
        switch (c1) {
            case 'S':
                b1 = true;
                break;
            case 'J':
                b1 = true;
                break;
            case '|':
                b1 = true;
                break;
            case 'L':
                b1 = true;
                break;
            default:
                break;
        }

        switch (c2) {
            case '7':
                b2 = true;
                break;
            case '|':
                b2 = true;
                break;
            case 'F':
                b2 = true;
                break;
            default:
                break;
        }

        return b1 && b2;
    }

    public static boolean checkDOWN(String s1, String s2) {
        char c1 = s1.charAt(0);
        char c2 = s2.charAt(0);

        boolean b1 = false;
        boolean b2 = false;
        switch (c1) {
            case 'S':
                b1 = true;
                break;
            case 'F':
                b1 = true;
                break;
            case '|':
                b1 = true;
                break;
            case '7':
                b1 = true;
                break;
            default:
                break;
        }

        switch (c2) {
            case 'J':
                b2 = true;
                break;
            case '|':
                b2 = true;
                break;
            case 'L':
                b2 = true;
                break;
            default:
                break;
        }

        return b1 && b2;
    }

    public static boolean checkRIGHT(String s1, String s2) {
        char c1 = s1.charAt(0);
        char c2 = s2.charAt(0);

        boolean b1 = false;
        boolean b2 = false;
        switch (c1) {
            case 'S':
                b1 = true;
                break;
            case 'F':
                b1 = true;
                break;
            case '-':
                b1 = true;
                break;
            case 'L':
                b1 = true;
                break;
            default:
                break;
        }

        switch (c2) {
            case 'J':
                b2 = true;
                break;
            case '-':
                b2 = true;
                break;
            case '7':
                b2 = true;
                break;
            default:
                break;
        }

        return b1 && b2;
    }

    public static boolean checkLEFT(String s1, String s2) {
        char c1 = s1.charAt(0);
        char c2 = s2.charAt(0);

        boolean b1 = false;
        boolean b2 = false;
        switch (c1) {
            case 'S':
                b1 = true;
                break;
            case '-':
                b1 = true;
                break;
            case 'J':
                b1 = true;
                break;
            case '7':
                b1 = true;
                break;
            default:
                break;
        }

        switch (c2) {
            case '-':
                b2 = true;
                break;
            case 'L':
                b2 = true;
                break;
            case 'F':
                b2 = true;
                break;
            default:
                break;
        }

        return b1 && b2;
    }

    public static void main(String[] args) {
        try {
            Scanner sc = new Scanner(new File("Input.txt"));

            int xSize = 10;
            int ySize = 10;

            if (args.length > 1) {
                xSize = Integer.parseInt(args[0]);
                ySize = Integer.parseInt(args[1]);
            }

            int xStarting = 0;
            int yStarting = 0;

            String[][] array = new String[xSize][ySize];
            boolean[][] alreadyChecked = new boolean[xSize][ySize];
            LinkedList<Pair> queue = new LinkedList<>();

            int index = 0;

            while (sc.hasNextLine()) {
                String[] data = sc.nextLine().split("");

                for (int i = 0; i < ySize; i++) {
                    array[index][i] = data[i];

                    if (data[i].equals("S")) {
                        xStarting = index;
                        yStarting = i;
                    }

                }
                index++;
            }

            queue.add(new Pair(xStarting, yStarting));
            alreadyChecked[xStarting][yStarting] = true;
            int seen = 1;
            LinkedList<Pair> points = new LinkedList<>();

            while (!queue.isEmpty()) {
                Pair p = queue.getFirst();
                queue.removeFirst();
                String s = array[p.x][p.y];

                if (array[p.x][p.y].equals("J") || array[p.x][p.y].equals("L") || array[p.x][p.y].equals("F")
                        || array[p.x][p.y].equals("7") || array[p.x][p.y].equals("S")) {
                    // System.out.println(array[p.x][p.y] + " x: " + p.x + " y: " + p.y);
                    points.add(new Pair(p.x, p.y));
                }

                if (0 < p.x && !alreadyChecked[p.x - 1][p.y] && checkUP(s, array[p.x - 1][p.y])) {
                    queue.add(new Pair(p.x - 1, p.y));
                    seen++;
                    alreadyChecked[p.x - 1][p.y] = true;
                    continue;
                }

                if (xSize - 1 > p.x && !alreadyChecked[p.x + 1][p.y] && checkDOWN(s, array[p.x + 1][p.y])) {
                    queue.add(new Pair(p.x + 1, p.y));
                    seen++;
                    alreadyChecked[p.x + 1][p.y] = true;
                    continue;
                }

                if (0 < p.y && !alreadyChecked[p.x][p.y - 1] && checkLEFT(s, array[p.x][p.y - 1])) {
                    queue.add(new Pair(p.x, p.y - 1));
                    seen++;
                    alreadyChecked[p.x][p.y - 1] = true;
                    continue;
                }

                if (ySize - 1 > p.y && !alreadyChecked[p.x][p.y + 1] && checkRIGHT(s, array[p.x][p.y + 1])) {
                    queue.add(new Pair(p.x, p.y + 1));
                    seen++;
                    alreadyChecked[p.x][p.y + 1] = true;
                }
            }

            /*
             * 
             * for (int i = 0; i < alreadyChecked.length; i++) {
             * for (int j = 0; j < alreadyChecked[i].length; j++) {
             * if (alreadyChecked[i][j]) {
             * System.out.print("# ");
             * } else {
             * System.out.print(". ");
             * }
             * }
             * System.out.println();
             * }
             * System.out.println("Seen: " + (seen / 2));
             */

            // FOR PART TWO
            points.addFirst(points.getLast());
            points.addLast(points.getFirst());
            int A = 0;
            for (int i = 1; i < points.size() - 1; i++) {
                A += points.get(i).x * (points.get(i - 1).y - points.get(i + 1).y);
            }

            A = Math.abs(A) / 2;
            // seen /= 2;
            int smallI = A - seen / 2 + 1;
            // smallI += seen;
            System.out.println(smallI);

            sc.close();
        } catch (Exception e) {
            System.err.println(e);
        }
    }
}