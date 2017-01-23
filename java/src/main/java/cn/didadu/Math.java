package cn.didadu;

/**
 * Created by zhangjing on 17-1-23.
 */
public class Math {
    public static void main(String[] args) {
        System.out.println(circumference(1.6f));
    }

    public static float circumference(float r) {
        float pi = 3.14f;
        //float c = 2 * pi * r;
        return  2 * pi * r;
    }
}
