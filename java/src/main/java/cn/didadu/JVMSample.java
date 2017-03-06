package cn.didadu;

/**
 * Created by zhangjing on 17-3-6.
 */
public class JVMSample {

    public static class TestExc extends RuntimeException {
    }

    void cantBeZero(int i) {
        if (i == 0) {
            throw new RuntimeException();
        }
    }

    void catchOne() {
        try {
            tryItOut();
        } catch (TestExc e) {
            handleExc(e);
        }
    }

    void tryItOut() {
        throw new TestExc();
    }

    void handleExc(RuntimeException exc) {

    }
}
