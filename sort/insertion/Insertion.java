public class Insertion {

    public static void main(String[] args) {
        Integer[] oo = new Integer[]{34, 8, 64, 51, 32, 21};
	    sort(oo);
        for (Integer var : oo) {
            System.out.println(var);
        }
    }

    public static <T> void sort(T[] array) {
        for (int i = 1; i< array.length; i++) {
            T it = array[i];
            for (int j = i; j > 0; j--) {
                T jt = array[j - 1];
                if (compare(jt, it) < 0) {
                    break;
                }else {
                    array[j] = jt;
                    array[j - 1] = it;
                }
            }
        }
    }

    public static <T> int compare(T o1, T o2) {
        return (Integer)o1 - (Integer)o2;
    }

}