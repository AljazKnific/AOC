import java.io.File;
import java.util.LinkedList;
import java.util.Scanner;

class Lens {
    String name;
    int val;

    public Lens(String name, int val) {
        this.name = name;
        this.val = val;

    }
}

public class D {

    public static void delete(LinkedList<LinkedList<Lens>> list, String data, int box) {
        // System.out.println("FROM " + box + " delete " + data);
        LinkedList<Lens> l = list.get(box);

        for (int i = 0; i < l.size(); i++) {
            if (l.get(i).name.equals(data)) {
                l.remove(i);
                return;
            }
        }
    }

    public static void addTo(LinkedList<LinkedList<Lens>> list, int box, String data, int val) {
        // System.out.println("FROM " + box + " addto " + data + " val " + val);
        LinkedList<Lens> l = list.get(box);

        for (int i = 0; i < l.size(); i++) {
            Lens le = l.get(i);
            if (le.name.equals(data)) {
                le.val = val;
                return;
            }
        }
        l.add(new Lens(data, val));
    }

    public static int getSum(LinkedList<LinkedList<Lens>> list) {
        int sum = 0;
        for (int i = 0; i < list.size(); i++) {
            for (int j = 0; j < list.get(i).size(); j++) {
                System.out
                        .println("BOX " + i + "j " + j + " " + list.get(i).get(j).name + " " + list.get(i).get(j).val);
                sum += ((i + 1) * (j + 1) * list.get(i).get(j).val);
            }
        }

        return sum;
    }

    public static void main(String[] args) {

        try {
            Scanner sc = new Scanner(new File("T.txt"));
            LinkedList<LinkedList<Lens>> list = new LinkedList<>();
            for (int i = 0; i < 256; i++) {
                list.add(new LinkedList<>());
            }

            String[] data = sc.nextLine().split(",");
            for (int i = 0; i < data.length; i++) {
                int res = 0;
                int j = 0;
                for (j = 0; j < data[i].length() && !(data[i].charAt(j) == '=' || data[i].charAt(j) == '-'); j++) {
                    int x = (int) data[i].charAt(j);
                    x += res;
                    x *= 17;
                    x %= 256;
                    res = x;
                }

                if (data[i].charAt(j) == '=') {
                    addTo(list, res, data[i].substring(0, j),
                            Integer.parseInt(data[i].substring(j + 1, data[i].length())));
                } else {
                    delete(list, data[i].substring(0, j), res);
                }

            }
            System.out.println("SUM: " + getSum(list));
        } catch (Exception e) {
            System.err.println(e);
        }
    }
}
