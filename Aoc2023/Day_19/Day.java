import java.io.File;
import java.util.Arrays;
import java.util.HashSet;
import java.util.LinkedList;
import java.util.Scanner;

class Inst {
    int num;
    int val;
    char sign;
    Info dest;
    String wholeIns;

    public Inst(int num, int val, char sign, Info dest, String wholeIns) {
        this.num = num;
        this.val = val;
        this.sign = sign;
        this.dest = dest;
        this.wholeIns = wholeIns;
    }

    public Inst(int num, int val, char sign, Info dest) {
        this.num = num;
        this.val = val;
        this.sign = sign;
        this.dest = dest;
    }
}

class Info {
    String name;
    LinkedList<Inst> instru;

    public Info(String name) {
        this.name = name;
        this.instru = new LinkedList<>();
    }
}

class Borders {
    int low;
    int high;
    int num;

    public Borders(int num) {
        this.num = num;
        this.low = 1;
        this.high = 4000;
    }

    public void setLow(int val) {
        this.low = Math.max(val, low) + 1;
    }

    public void setHigh(int val) {
        this.high = Math.min(val, high) - 1;
    }

}

public class Day {

    public static Info getInfo(String name, HashSet<Info> set) {
        for (Info info : set) {
            if (info.name.equals(name)) {
                return info;
            }
        }
        Info in = new Info(name);
        set.add(in);
        return in;
    }

    public static Info addToSet(String data, HashSet<Info> set) {
        String[] temp = data.split("\\{");
        Info info = getInfo(temp[0], set);
        temp = temp[1].split(",");

        for (int i = 0; i < temp.length - 1; i++) {
            String[] arr = temp[i].split(":");
            // System.out.println(arr[0]);

            int num = -1;
            int val = Integer.parseInt(arr[0].substring(2, arr[0].length()));
            char sign = arr[0].charAt(1);
            Info dest = getInfo(arr[1], set);
            switch (arr[0].charAt(0)) {
                case 'x':
                    num = 0;
                    break;
                case 'm':
                    num = 1;
                    break;
                case 'a':
                    num = 2;
                    break;
                case 's':
                    num = 3;
                    break;

                default:
                    System.out.println("Wrong number");
                    break;
            }
            info.instru.add(new Inst(num, val, sign, dest, arr[0]));

        }
        // we get the last one
        Info dest = getInfo(temp[temp.length - 1], set);
        info.instru.add(new Inst(0, -42, '>', dest, "X"));

        return info;

    }

    public static boolean statement(int a, int b, char op) {
        if (op == '>') {
            return a > b;
        } else {
            return a < b;
        }
    }

    public static int calc(LinkedList<Integer> list, Info start) {

        while (!(start.name.equals("A") || start.name.equals("R"))) {
            for (Inst inst : start.instru) {
                if (statement(list.get(inst.num), inst.val, inst.sign)) {
                    start = inst.dest;
                    break;
                }
            }
        }
        if (start.name.equals("A")) {
            return list.get(0) + list.get(1) + list.get(2) + list.get(3);
        } else {
            return 0;
        }
    }

    public static String getCorrectForm(String s) {
        if (s.length() < 2) {
            return "";
        }

        if (s.charAt(1) == '>') {
            return s.substring(0, 1) + "<" + s.substring(2);
        } else {
            return s.substring(0, 1) + ">" + s.substring(2);
        }
    }

    public static void calcBorder(String data, LinkedList<Borders> bor) {
        int b = -1;
        switch (data.charAt(0)) {
            case 'x':
                b = 0;
                break;
            case 'm':
                b = 1;
                break;
            case 'a':
                b = 2;
                break;
            case 's':
                b = 3;
                break;
            default:
                System.err.println("WRONG INPUT");
                break;
        }

        Borders temp = bor.get(b);

        if (data.charAt(1) == '>') {
            temp.setLow(Integer.parseInt(data.substring(2)));
        } else {
            temp.setHigh(Integer.parseInt(data.substring(2)));
        }
    }

    public static void paths(Info in, String s, LinkedList<LinkedList<Borders>> partTwo) {
        if (in.name.equals("A")) {
            System.out.println(s);
            String[] arr = s.split(",");
            // System.out.println(Arrays.toString(arr));
            LinkedList<Borders> b = new LinkedList<>();
            b.add(new Borders(0));
            b.add(new Borders(1));
            b.add(new Borders(2));
            b.add(new Borders(3));
            for (int i = 1; i < arr.length; i++) {
                calcBorder(arr[i], b);
            }
            partTwo.add(b);
            return;
        }

        if (in.name.equals("R")) {
            return;
        }
        // add also the instructions that werent added
        for (int i = 0; i < in.instru.size(); i++) {
            if (i == in.instru.size() - 1) {
                paths(in.instru.get(i).dest, s, partTwo);
            } else {
                paths(in.instru.get(i).dest, s + "," + in.instru.get(i).wholeIns, partTwo);
                s += "," + getCorrectForm(in.instru.get(i).wholeIns);
            }
        }

    }

    public static void main(String[] args) {

        try {

            Scanner sc = new Scanner(new File("I.txt"));

            HashSet<Info> set = new HashSet<>();
            set.add(new Info("R"));
            set.add(new Info("A"));

            boolean first = true;
            Info start = null;
            int num = 0;

            while (sc.hasNextLine()) {
                String data = sc.nextLine();
                if (data.length() == 0) {
                    first = false;
                    continue;
                }

                if (first) {
                    data = data.substring(0, data.length() - 1);
                    if (data.substring(0, 2).equals("in")) {
                        start = addToSet(data, set);
                    } else {
                        addToSet(data, set);

                    }

                } else {
                    data = data.substring(1, data.length() - 1);
                    String[] arr = data.split(",");
                    LinkedList<Integer> nums = new LinkedList<>();
                    for (int i = 0; i < arr.length; i++) {
                        String[] temp = arr[i].split("=");
                        nums.add(Integer.parseInt(temp[1]));

                    }

                    num += calc(nums, start);

                }
            }

            System.out.println(num);
            /*
             * 
             * for (Info info : set) {
             * System.out.println(info.name);
             * for (Inst info2 : info.instru) {
             * System.out.println(info2.wholeIns);
             * }
             * }
             */
            LinkedList<LinkedList<Borders>> partTwo = new LinkedList<>();
            paths(start, "", partTwo);
            Long res = 0L;
            for (LinkedList<Borders> linkedList : partTwo) {
                Long x = 1L;
                for (int i = 0; i < linkedList.size(); i++) {
                    Borders tt = linkedList.get(i);

                    x *= (tt.high - tt.low + 1);
                    System.out.print(
                            "| low: " + tt.low + " high: " + tt.high + " -> " + " , x: " + (tt.high - tt.low + 1));
                }
                res += x;
                System.out.println();
            }

            System.out.println(res);

            // IDEA
            // For part two find all possible solutions, that reach string A.
            // After that for each sequence of rules calculate the intervals for each number
            // Then we just multiply these numbers and

            sc.close();
        } catch (Exception e) {
            System.err.println(e);
        }

    }
}
