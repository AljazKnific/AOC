import java.io.File;
import java.util.Arrays;
import java.util.LinkedList;
import java.util.Scanner;

public class Day {

    public static int rows(String[][] arr, int x, int y, boolean trans) {
        int sum = 0;
        /*
         * 
         * for (int i = 0; i < x; i++) {
         * for (int j = 0; j < y; j++) {
         * System.out.print(arr[i][j]);
         * }
         * System.out.println();
         * }
         */

        // System.out.println("START-> X: " + x + " Y: " + y);

        for (int i = 1; i < x; i++) {
            int smudge = 0;
            int smaller = i;
            boolean bo = true;
            if (x - i < i) {
                smaller = x - i;
            }

            // System.out.println("Smaller -> " + smaller);

            // System.out.println(i + " : " + x);

            for (int j = 0; j < smaller; j++) {
                for (int k = 0; k < y; k++) {
                    // System.out.println(arr[j + i][k] + " " + arr[i - j - 1][k]);
                    if (!arr[j + i][k].equals(arr[i - j - 1][k])) {
                        smudge++;

                        if (smudge == 2) {
                            bo = false;
                            break;
                        }
                    }
                }
                System.out.println("I -> " + i + " smudge: " + smudge);
                if (!bo) {
                    break;
                }
            }

            if (bo && smudge == 1) {
                if (trans) {
                    return i;
                }
                return i * 100;
            }

        }

        return sum;

    }

    public static String[][] arT(String[][] arr, int x, int y) {

        String[][] temp = new String[y][x];

        for (int i = 0; i < x; i++) {
            for (int j = 0; j < y; j++) {
                temp[j][i] = arr[i][j];
            }
        }
        /*
         * 
         * for (int i = 0; i < y; i++) {
         * for (int j = 0; j < x; j++) {
         * System.out.print(temp[i][j]);
         * }
         * System.out.println();
         * }
         */

        return temp;
    }

    public static void main(String[] args) {
        try {
            Scanner sc = new Scanner(new File("Input.txt"));
            String[][] array = new String[200][200];

            int index = 0;
            int xS = 0;
            int sum = 0;
            while (sc.hasNextLine()) {
                String[] data = sc.nextLine().split("");
                if (data.length > 1) {
                    xS = data.length;
                }
                array[index] = data;

                if (data.length == 1) {
                    // call the function
                    sum += rows(array, index, xS, false);
                    // System.out.println(sum);
                    // System.out.println("End of first");
                    sum += rows(arT(array, index, xS), xS, index, true);
                    // System.out.println("End of T");
                    index = -1;

                }

                index++;

            }

            sum += rows(array, index, xS, false);
            // System.out.println("End of first");
            sum += rows(arT(array, index, xS), xS, index, true);

            System.out.println("SUM: " + sum);

            sc.close();

        } catch (Exception e) {
            System.err.println(e);
        }
    }

}
