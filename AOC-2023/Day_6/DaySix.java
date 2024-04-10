import java.io.File;
import java.util.*;

class DaySix {

    public static int partOne(int t, int d) {
        // quadratic equation x^2 - tx + d = 0

        double highroot = (t + Math.pow((double) (t * t) - (4 * d), 0.5)) / 2.0;
        double lowroot = (t - Math.pow((double) (t * t) - (4 * d), 0.5)) / 2.0;

        System.out.println("T: " + t + " D: " + d + " HR: " + highroot + " LR: " + lowroot);

        return (int) (Math.floor(highroot) - Math.ceil(lowroot) + 1);
    }

    public static int partTwo(int t, Long d) {
        // quadratic equation x^2 - tx + d = 0

        double highroot = (t + Math.pow((double) (t * t) - (4 * d), 0.5)) / 2.0;
        double lowroot = (t - Math.pow((double) (t * t) - (4 * d), 0.5)) / 2.0;

        System.out.println("T: " + t + " D: " + d + " HR: " + highroot + " LR: " + lowroot);

        return (int) (Math.floor(highroot) - Math.ceil(lowroot) + 1);
    }

    public static void main(String[] args) {
        try {
            Scanner sc = new Scanner(new File("Input.txt"));

            LinkedList<Integer> time = new LinkedList<>();
            LinkedList<Integer> dist = new LinkedList<>();

            List<String> data = new LinkedList<String>(Arrays.asList(sc.nextLine().split(":")[1].trim().split(" ")));
            data.removeAll(Arrays.asList("", null));
            Long timeMain = 0L;

            for (String string : data) {
                timeMain *= (int) Math.pow(10, string.length());
                timeMain += (Integer.parseInt(string));
                time.add(Integer.parseInt(string));
            }

            data = new LinkedList<String>(Arrays.asList(sc.nextLine().split(":")[1].trim().split(" ")));
            data.removeAll(Arrays.asList("", null));
            Long distMain = 0L;

            for (String string : data) {
                distMain *= (int) Math.pow(10, string.length());
                distMain += (Integer.parseInt(string));
                dist.add(Integer.parseInt(string));
            }

            int sum = 1;
            int sum2 = 0;

            for (int i = 0; i < time.size(); i++) {
                int t = time.get(i);
                int d = dist.get(i);
                int num = 0;

                for (int j = 0; j <= t; j++) {
                    if (j * (t - j) > d) {
                        num++;
                    }
                }
                System.out.println("NUM: " + num);
                sum *= num;
                // sum *= partOne(t, d);

            }

            for (int i = 0; i < timeMain; i++) {
                if (i * (timeMain - i) > distMain) {
                    sum2++;
                }
            }

            // sum2 = partTwo(timeMain, distMain);
            System.out.println("Res: " + sum);
            System.out.println("Res2: " + sum2);

        } catch (Exception e) {
            System.err.println(e);
        }
    }
}