import java.io.File;
import java.util.Arrays;
import java.util.Scanner;

class DayTwo {

    public static boolean eval(String[] arr, int red, int green, int blue) {

        arr = arr[1].split(";");

        for (int i = 0; i < arr.length; i++) {
            String[] temp = arr[i].split(",");

            for (int j = 0; j < temp.length; j++) {
                String[] t = temp[j].split(" ");

                for (int k = 1; k < t.length; k += 2) {

                    int num = Integer.parseInt(t[k]);

                    switch (t[k + 1]) {

                        case "red":
                            if (num > red) {
                                return false;
                            }
                            break;

                        case "blue":
                            if (num > blue) {
                                return false;
                            }
                            break;

                        case "green":
                            if (num > green) {
                                return false;
                            }
                            break;

                        default:
                            System.out.println("Failed, wrong input");
                            break;
                    }
                }

            }

        }

        return true;
    }

    public static int eval2(String[] arr) {

        int r = 0;
        int g = 0;
        int b = 0;

        arr = arr[1].split(";");

        for (int i = 0; i < arr.length; i++) {
            String[] temp = arr[i].split(",");

            for (int j = 0; j < temp.length; j++) {
                String[] t = temp[j].split(" ");

                for (int k = 1; k < t.length; k += 2) {

                    int num = Integer.parseInt(t[k]);

                    switch (t[k + 1]) {

                        case "red":
                            if (r < num)
                                r = num;
                            break;

                        case "blue":
                            if (b < num)
                                b = num;

                            break;

                        case "green":
                            if (g < num)
                                g = num;

                            break;

                        default:
                            System.out.println("Failed, wrong input");
                            break;
                    }
                }

            }

        }

        return g * r * b;
    }

    public static void main(String[] args) {

        try {

            File inp = new File("Input.txt");
            Scanner sc = new Scanner(inp);

            int sum = 0;
            int sumTwo = 0;

            while (sc.hasNextLine()) {
                String data = sc.nextLine();
                String[] arr = data.split(":");

                int game = Integer.valueOf(arr[0].split(" ")[1]);

                if (eval(arr, 12, 13, 14)) {
                    sum += game;
                }

                sumTwo += eval2(arr);
            }

            System.out.println("Sum: " + sum);
            System.out.println("Sum2: " + sumTwo);

            sc.close();

        } catch (Exception e) {
            System.err.println(e);
        }

    }
}