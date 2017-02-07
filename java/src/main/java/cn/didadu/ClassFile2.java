package cn.didadu;

/**
 * 测试CONSTANT_NameAndType_info
 * Created by zhangjing on 17-1-11.
 */
public class ClassFile2 {
    public static final boolean FLAG = true;
    public static final byte BYTE = 123;
    public static final char X = 'X';
    public static final short SHORT = 12345;
    public static final int INT = 123456789;
    public static final long LONG = 12345678901L;
    public static final float PI = 3.14f;
    public static final double E = 2.71828;

    public static void main(String[] args) throws RuntimeException {
        /**
         * 测试局部变量在class文件中出现的位置
         */
        String localStr1 = "Hello,";
        String localStr2 = "World";
        // 下面字符串的+编译的时候会被优化成StringBuilder
        System.out.println(localStr1 + localStr2);
    }
}
