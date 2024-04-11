import java.io.File;
import java.util.Arrays;
import java.util.LinkedList;
import java.util.Scanner;

public class DayNine {

    public static int partOne(LinkedList<Integer> list) {
        boolean zero = true;
        LinkedList<Integer> resList = new LinkedList<>();

        for (int i = 0; i < list.size() - 1; i++) {
            Integer fir = list.get(i);
            Integer sec = list.get(i + 1);
            // check for zeroes
            zero &= (fir == 0);
            zero &= (sec == 0);

            resList.add(sec - fir);

        }

        if (zero)
            return 0;

        return (list.getLast() + partOne(resList));
    }

    public static int partTwo(LinkedList<Integer> list) {
        boolean zero = true;
        LinkedList<Integer> resList = new LinkedList<>();

        for (int i = 0; i < list.size() - 1; i++) {
            Integer fir = list.get(i);
            Integer sec = list.get(i + 1);
            // check for zeroes
            zero &= (fir == 0);
            zero &= (sec == 0);

            resList.add(sec - fir);

        }

        if (zero)
            return 0;

        return (list.getFirst() - partTwo(resList));
    }

    public static void main(String[] args) {

        try {
            Scanner sc = new Scanner(new File("Input.txt"));

            int sum = 0;
            int sum2 = 0;

            while (sc.hasNextLine()) {
                String[] data = sc.nextLine().split(" ");
                LinkedList<Integer> list = new LinkedList<>();

                for (String string : data) {
                    list.add(Integer.parseInt(string));
                }
                // partOne
                sum += partOne(list);

                // partTwo
                sum2 += partTwo(list);
            }

            System.out.println("Sum: " + sum);
            System.out.println("SUM2: " + sum2);
            sc.close();

        } catch (Exception e) {
            System.err.println(e);
        }
    }

}
