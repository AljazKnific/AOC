import java.io.File;
import java.util.Arrays;
import java.util.HashMap;
import java.util.LinkedList;
import java.util.Scanner;

class Pair {
    int start;
    int end;
    boolean bool;

    public Pair(int x, int y) {
        start = x;
        end = y;
        bool = true;
    }

    public void setStart(int x) {
        start = x;
    }

    public void setEnd(int x) {
        end = x;
    }

    public int getStart() {
        return start;
    }

    public int getEnd() {
        return end;
    }

    public boolean getBool() {
        return bool;
    }

    public void toggleBool() {
        bool = !bool;
    }
}

class Res {
    int val = 1;
    int yes = 0;
}

class Group {

    LinkedList<Integer> stars = new LinkedList<>();
    HashMap<Integer, LinkedList<Pair>> numbers = new HashMap<>();

    public Group() {

    }

    public void addStar(int x) {
        stars.add(x);
    }

    public LinkedList<Integer> getStars() {
        return stars;
    }

    public HashMap<Integer, LinkedList<Pair>> getNumbers() {
        return numbers;
    }

    public void addNumber(int val, int start, int end) {
        Pair p = new Pair(start, end);
        LinkedList<Pair> temp = numbers.get(val);
        if (temp != null) {
            temp.add(p);
        } else {
            LinkedList<Pair> nL = new LinkedList<>();
            nL.add(p);
            numbers.put(val, nL);
        }

    }

    public void izpis() {
        System.out.println("Zvezdice:");
        for (int i = 0; i < stars.size(); i++) {
            System.out.print(stars.get(i) + " ");
        }
        System.out.println();

        System.out.println("Stevilke: velikost: " + numbers.size());
        for (Integer integer : numbers.keySet()) {
            LinkedList<Pair> t = numbers.get(integer);
            for (Pair pair : t) {
                System.out.println("Start: " + pair.getStart() + " End: " + pair.getEnd() + " -> " + integer);
            }
        }
    }
}

class DayThree {

    public static Res sumAroundX(Group g, int x) {
        Res vrni = new Res();
        int sum = 0;
        for (Integer t : g.getNumbers().keySet()) {
            LinkedList<Pair> temp = g.getNumbers().get(t);
            for (Pair p : temp) {
                if (p.getStart() <= x && x <= p.getEnd() && p.getBool()) {
                    sum += t;
                    vrni.val = t;
                    vrni.yes += 1;
                    p.toggleBool();

                    return vrni;
                }
            }
        }

        return vrni;
    }

    public static int multi(LinkedList<Res> r) {
        int mul = 1;
        int amount = 0;
        for (Res res : r) {

            if (res.yes == 1) {
                amount++;
                mul *= res.val;
            }
        }
        if (amount == 2)
            return mul;
        return 0;
    }

    public static int sumAround(LinkedList<Group> list, int x, int y, int maxX) {
        int sum = 0;

        int maxY = list.size();
        LinkedList<Res> rLinkedList = new LinkedList<>();

        if (y == 0) {
            // lev zgornji kot
            if (x == 0) {
                Res r1 = sumAroundX(list.get(0), 0);
                Res r2 = sumAroundX(list.get(1), 1);
                Res r3 = sumAroundX(list.get(1), 0);

                rLinkedList.add(r1);
                rLinkedList.add(r2);
                rLinkedList.add(r3);

                sum += multi(rLinkedList);

                // desen zgornji kot
            } else if (x == maxX) {

                Res r1 = sumAroundX(list.get(0), maxX - 1);
                Res r2 = sumAroundX(list.get(1), maxX);
                Res r3 = sumAroundX(list.get(1), maxX - 1);

                rLinkedList.add(r1);
                rLinkedList.add(r2);
                rLinkedList.add(r3);

                sum += multi(rLinkedList);

                // sredina navrh
            } else {
                Res r1 = sumAroundX(list.get(0), x + 1);
                Res r2 = sumAroundX(list.get(0), x - 1);
                Res r3 = sumAroundX(list.get(1), x + 1);
                Res r4 = sumAroundX(list.get(1), x);
                Res r5 = sumAroundX(list.get(1), x - 1);

                rLinkedList.add(r1);
                rLinkedList.add(r2);
                rLinkedList.add(r3);
                rLinkedList.add(r4);
                rLinkedList.add(r5);

                sum += multi(rLinkedList);
            }

        } else if (y == maxY - 1) {
            // lev spodnji kot
            if (x == 0) {
                Res r1 = sumAroundX(list.get(maxY - 1), 1);
                Res r2 = sumAroundX(list.get(maxY - 2), 1);
                Res r3 = sumAroundX(list.get(maxY - 2), 0);

                rLinkedList.add(r1);
                rLinkedList.add(r2);
                rLinkedList.add(r3);

                sum += multi(rLinkedList);
                // desen spodjni kot
            } else if (x == maxX) {
                Res r1 = sumAroundX(list.get(maxY - 1), maxX - 1);
                Res r2 = sumAroundX(list.get(maxY - 2), maxX);
                Res r3 = sumAroundX(list.get(maxY - 2), maxX - 1);

                rLinkedList.add(r1);
                rLinkedList.add(r2);
                rLinkedList.add(r3);

                sum += multi(rLinkedList);
                // spodaj sredina
            } else {
                Res r1 = sumAroundX(list.get(maxY - 1), x + 1);
                Res r2 = sumAroundX(list.get(maxY - 1), x - 1);
                Res r3 = sumAroundX(list.get(maxY - 2), x + 1);
                Res r4 = sumAroundX(list.get(maxY - 2), x);
                Res r5 = sumAroundX(list.get(maxY - 2), x - 1);

                rLinkedList.add(r1);
                rLinkedList.add(r2);
                rLinkedList.add(r3);
                rLinkedList.add(r4);
                rLinkedList.add(r5);

                sum += multi(rLinkedList);
            }
        } else {
            if (x == 0) {
                Res r1 = sumAroundX(list.get(y), x + 1);
                Res r2 = sumAroundX(list.get(y - 1), x + 1);
                Res r3 = sumAroundX(list.get(y + 1), x + 1);
                Res r4 = sumAroundX(list.get(y - 1), x);
                Res r5 = sumAroundX(list.get(y + 1), x);

                rLinkedList.add(r1);
                rLinkedList.add(r2);
                rLinkedList.add(r3);
                rLinkedList.add(r4);
                rLinkedList.add(r5);

                sum += multi(rLinkedList);
            } else if (x == maxX) {
                Res r1 = sumAroundX(list.get(y), maxX - 1);
                Res r2 = sumAroundX(list.get(y - 1), maxX - 1);
                Res r3 = sumAroundX(list.get(y + 1), maxX - 1);
                Res r4 = sumAroundX(list.get(y - 1), maxX);
                Res r5 = sumAroundX(list.get(y + 1), maxX);

                rLinkedList.add(r1);
                rLinkedList.add(r2);
                rLinkedList.add(r3);
                rLinkedList.add(r4);
                rLinkedList.add(r5);

                sum += multi(rLinkedList);
            } else {
                Res r1 = sumAroundX(list.get(y), x + 1);
                Res r2 = sumAroundX(list.get(y), x - 1);
                Res r3 = sumAroundX(list.get(y + 1), x + 1);
                Res r4 = sumAroundX(list.get(y + 1), x);
                Res r5 = sumAroundX(list.get(y + 1), x - 1);
                Res r6 = sumAroundX(list.get(y - 1), x + 1);
                Res r7 = sumAroundX(list.get(y - 1), x);
                Res r8 = sumAroundX(list.get(y - 1), x - 1);

                rLinkedList.add(r1);
                rLinkedList.add(r2);
                rLinkedList.add(r3);
                rLinkedList.add(r4);
                rLinkedList.add(r5);
                rLinkedList.add(r6);
                rLinkedList.add(r7);
                rLinkedList.add(r8);

                sum += multi(rLinkedList);
            }
        }

        return sum;
    }

    public static void main(String[] args) {

        try {
            File inp = new File("Input.txt");
            Scanner sc = new Scanner(inp);

            LinkedList<Group> list = new LinkedList<>();

            int sum = 0;
            int x = 0;

            while (sc.hasNextLine()) {
                String data = sc.nextLine();
                x = data.length();
                Group g = new Group();

                for (int i = 0; i < data.length(); i++) {
                    if (data.charAt(i) == '.') {
                        continue;
                    }

                    // pregledamo zacetek in konec ter vrednost inta
                    if (Character.isDigit(data.charAt(i))) {
                        int start = i;

                        int num = 0;

                        while (i != data.length() && Character.isDigit(data.charAt(i))) {
                            num *= 10;
                            num += Character.getNumericValue(data.charAt(i));
                            i++;
                        }
                        i--;

                        g.addNumber(num, start, i);

                    } else if (data.charAt(i) == '*') {
                        // dodamo index zvezdice
                        g.addStar(i);
                    }
                }

                list.add(g);
            }

            for (int i = 0; i < list.size(); i++) {
                // list.get(i).izpis();
                Group g = list.get(i);

                for (int j = 0; j < g.getStars().size(); j++) {
                    int a = g.getStars().get(j);
                    // a == x-os, i == y-os
                    sum += sumAround(list, a, i, x - 1);
                }

            }

            System.out.println(sum);

            sc.close();

        } catch (Exception e) {
            System.err.println(e);
        }

    }
}