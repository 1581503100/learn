package lsj.cc.flag;

/**
 * @Author:sjliu7
 * @Date:2018/11/14 15:47
 */
public class StringFlag {
    private String name;
    private String def;
    private String usage;

    public String getName() {
        return name;
    }

    public StringFlag name(String name) {
        this.name = name;
        return this;
    }

    public String getDef() {
        return def;
    }

    public StringFlag value(String value) {
        this.def = value;
        return this;
    }

    public String getUsage() {
        return usage;
    }

    public StringFlag usage(String usage) {
        this.usage = usage;
        return this;
    }
}
