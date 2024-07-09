package leetcode

import (
	"bufio"
	"container/list"
	"fmt"
	"math"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func main() {
	HJ30字符串合并处理()
}

// 给定两个字符串形式的非负整数 num1和num2 ，计算它们的和并同样以字符串形式返回。
// 你不能使用任何內建的用于处理大整数的库（比如 BigInteger）， 也不能直接将输入的字符串转换为整数形式。
// 示例 1：
// 输入：num1 = "11", num2 = "123"
// 输出："134"
// 示例 2：
// 输入：num1 = "456", num2 = "77"
// 输出："533"
func 字符串求和() {
	var num1, num2 string
	fmt.Scan(&num1, &num2)
	num1Len := len(num1)
	num2Len := len(num2)
	maxLen := max(num1Len, num2Len)
	if num1Len == maxLen {
		n := maxLen - num2Len
		for i := 0; i < n; i++ {
			num2 = "0" + num2
		}
	} else {
		n := maxLen - num1Len
		for i := 0; i < n; i++ {
			num1 = "0" + num1
		}
	}
	resStr := ""
	j := 0
	for i := maxLen - 1; i >= 0; i-- {
		n1, _ := strconv.Atoi(string(num1[i]))
		n2, _ := strconv.Atoi(string(num2[i]))
		n := n1 + n2 + j
		m := n - 10
		if m >= 0 {
			j = 1
		} else {
			j = 0
			m = n
		}
		resStr = strconv.FormatInt(int64(m), 10) + resStr
	}
	if j == 1 {
		resStr = "1" + resStr
	}
	fmt.Println(resStr)
}

// 计算字符串最后一个单词的长度
// 单词以空格隔开，字符串长度小于5000。（注：字符串末尾不以空格为结尾）
// 输入描述：
// 输入一行，代表要计算的字符串，非空，长度小于5000。
// 输出描述：
// 输出一个整数，表示输入字符串最后一个单词的长度。
func 计算字符串最后一个单词的长度() {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}
	input = strings.TrimSuffix(input, "\n")
	slice := strings.Split(input, " ")
	str := slice[len(slice)-1]
	fmt.Println(len(str))
}

// 计算某字符出现次数
// 写出一个程序，接受一个由字母、数字和空格组成的字符串，和一个字符，然后输出输入字符串中该字符的出现次数。（不区分大小写字母）
// 数据范围：
// 1≤n≤1000
// 输入描述：
// 第一行输入一个由字母、数字和空格组成的字符串，第二行输入一个字符（保证该字符不为空格）。
// 输出描述：
// 输出输入字符串中含有该字符的个数。（不区分大小写字母）
func 计算某字符出现次数() {
	reader := bufio.NewReader(os.Stdin)
	inputStr, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("err:%v\n", err)
	}
	str, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("err:%v\n", err)
	}
	inputStr = strings.TrimSuffix(inputStr, "\n")
	str = strings.TrimSuffix(str, "\n")
	inputStr = strings.ToLower(inputStr)
	str = strings.ToLower(str)

	count := strings.Count(inputStr, str)
	fmt.Println(count)
}

// 明明的随机数
// 明明生成了N个1到500之间的随机整数。请你删去其中重复的数字，即相同的数字只保留一个，把其余相同的数去掉，然后再把这些数从小到大排序，按照排好的顺序输出。
// 数据范围：1≤n≤1000  ，输入的数字大小满足1≤val≤500
// 输入描述：
// 第一行先输入随机整数的个数 N 。 接下来的 N 行每行输入一个整数，代表明明生成的随机数。 具体格式可以参考下面的"示例"。
// 输出描述：
// 输出多行，表示输入数据处理后的结果
func 明明的随机数() {
	reader := bufio.NewReader(os.Stdin)
	nstr, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("err:%v\n", err)
		return
	}
	nstr = strings.TrimSuffix(nstr, "\n")
	n, err := strconv.Atoi(nstr)
	if err != nil {
		fmt.Printf("err:%v\n", err)
		return
	}
	var randMap = make(map[int]int)
	for i := 0; i < n; i++ {
		randStr, _ := reader.ReadString('\n')
		randStr = strings.TrimSuffix(randStr, "\n")
		randInt, _ := strconv.Atoi(randStr)
		randMap[randInt] = randInt
	}
	var randSlic = make([]int, 0)
	for _, val := range randMap {
		randSlic = append(randSlic, val)
	}
	sort.Ints(randSlic)
	for _, val := range randSlic {
		fmt.Println(val)
	}
}

// 字符串分割
// •输入一个字符串，请按长度为8拆分每个输入字符串并进行输出；
// •长度不是8整数倍的字符串请在后面补数字0，空字符串不处理。
// 输入描述：
// 连续输入字符串(每个字符串长度小于等于100)
// 输出描述：
// 依次输出所有分割后的长度为8的新字符串
func 字符串分割() {
	reader := bufio.NewReader(os.Stdin)
	str, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("err:%v\n", str)
		return
	}
	str = strings.TrimSuffix(str, "\n")
	var n = 0
	strSlice := []string{}
	s := ""
	for _, val := range str {
		s = s + string(val)
		n++
		if n == 8 {
			strSlice = append(strSlice, s)
			s = ""
			n = 0
		}
	}
	if n != 0 {
		for i := 8 - n; i > 0; i-- {
			s = fmt.Sprintf("%s%s", s, "0")
		}
		strSlice = append(strSlice, s)
	}
	for _, val := range strSlice {
		fmt.Println(val)
	}
}

// 进制转换
// 写出一个程序，接受一个十六进制的数，输出该数值的十进制表示。
// 输入描述：
// 输入一个十六进制的数值字符串。
// 输出描述：
// 输出该数值的十进制字符串。不同组的测试用例用\n隔开。
func 进制转换1() {
	reader := bufio.NewReader(os.Stdin)
	str, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	str = strings.TrimSuffix(str, "\n")
	str = strings.TrimLeft(str, "0x")
	strSlice := strings.Split(str, "")
	for i := 0; i < len(strSlice)/2; i++ {
		strSlice[i], strSlice[len(strSlice)-1-i] = strSlice[len(strSlice)-1-i], strSlice[i]
	}

	sNumMap := map[string]int{
		"0": 0,
		"1": 1,
		"2": 2,
		"3": 3,
		"4": 4,
		"5": 5,
		"6": 6,
		"7": 7,
		"8": 8,
		"9": 9,
		"A": 10,
		"B": 11,
		"C": 12,
		"D": 13,
		"E": 14,
		"F": 15,
	}
	num := 0
	for k, v := range strSlice {
		n, ok := sNumMap[v]
		if !ok {
			fmt.Printf("输入的不是十六进制数:%s\n", str)
			return
		}
		num = num + n*(int(math.Pow(16, float64(k))))
	}
	fmt.Println(num)
}

// strconv.ParseInt()	字符串转换为整形
// strconv.FormatInt()	整形转换为字符串
// strconv.ParseFloat()	字符串转换为浮点型
// strconv.FormatFloat()浮点型转换为字符串
func 进制转换2() {
	var str string
	_, err := fmt.Scanln(&str)
	if err != nil {
		fmt.Printf("err:%v\n", err)
		return
	}
	str = strings.TrimPrefix(str, "0x")
	num, err := strconv.ParseInt(str, 16, 64)
	if err != nil {
		fmt.Printf("err:%v\n", err)
	}
	fmt.Println(num)
}

// 质数因子
// 功能:输入一个正整数，按照从小到大的顺序输出它的所有质因子（重复的也要列举）（如180的质因子为2 2 3 3 5 ）
// 输入描述：
// 输入一个整数
// 输出描述：
// 按照从小到大的顺序输出它的所有质数的因子，以空格隔开。
func 质数因子() {
	reader := bufio.NewReader(os.Stdin)
	str, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("err:%v\n", err)
		return
	}
	str = strings.TrimSuffix(str, "\n")
	num, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		fmt.Printf("err:%v\n", err)
		return
	}
	n := []int{}
	for i := 2; i*i <= int(num); i++ {
		for int(num)%i == 0 {
			num = int64(int(num) / i)
			n = append(n, i)
		}
	}
	if num != 1 {
		n = append(n, int(num))
	}

	sort.Ints(n)
	for _, v := range n {
		fmt.Printf("%d ", v)
	}
}

// 取近似值
// 写出一个程序，接受一个正浮点数值，输出该数值的近似整数值。如果小数点后数值大于等于 0.5 ,向上取整；小于 0.5 ，则向下取整。
// 数据范围：保证输入的数字在 32 位浮点数范围内
// 输入描述：
// 输入一个正浮点数值
// 输出描述：
// 输出该数值的近似整数值
func 取近似值1() {
	reader := bufio.NewReader(os.Stdin)
	str, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("err:%v\n", err)
		return
	}
	str = strings.TrimSuffix(str, "\n")
	strSlice := strings.Split(str, ".")
	nStr := strSlice[0]
	fStr := "0." + strSlice[1]
	nInt, err := strconv.ParseInt(nStr, 10, 64)
	if err != nil {
		fmt.Printf("err:%v\n", err)
		return
	}
	fFloat, err := strconv.ParseFloat(fStr, 64)
	if err != nil {
		fmt.Printf("err:%v\n", err)
		return
	}
	if fFloat >= 0.5 {
		nInt++
	}
	fmt.Println(nInt)
}

// math.Round()	取近似值
// math.Floor()	向下取整
// math.Ceil()	向上取整
func 取近似值2() {
	var str string
	_, err := fmt.Scanln(&str)
	if err != nil {
		fmt.Printf("err:%v\n", err)
		return
	}
	f, err := strconv.ParseFloat(str, 64)
	if err != nil {
		fmt.Printf("err:%v\n", err)
		return
	}

	//取近似值
	n := math.Round(f)
	fmt.Println(n)
}

// 合并表记录
// 数据表记录包含表索引index和数值value（int范围的正整数），请对表索引相同的记录进行合并，即将相同索引的数值进行求和运算，输出按照index值升序进行输出。
// 提示:
// 0 <= index <= 11111111
// 1 <= value <= 100000
// 输入描述：
// 先输入键值对的个数n（1 <= n <= 500）
// 接下来n行每行输入成对的index和value值，以空格隔开
// 输出描述：
// 输出合并后的键值对（多行）
func 合并表记录() {
	var n int
	_, err := fmt.Scanln(&n)
	if err != nil {
		fmt.Printf("err:%v\n", err)
		return
	}

	var k, v int
	keySlice := []int{}
	resultMap := map[int]int{}
	for i := 0; i < n; i++ {
		fmt.Scan(&k, &v)
		if value, ok := resultMap[k]; ok {
			resultMap[k] = value + v
		} else {
			resultMap[k] = v
		}
	}
	for key, _ := range resultMap {
		keySlice = append(keySlice, key)
	}
	sort.Ints(keySlice)
	for _, value := range keySlice {
		fmt.Printf("%d %d\n", value, resultMap[value])
	}
}

// 提取不重复的整数
// 输入一个 int 型整数，按照从右向左的阅读顺序，返回一个不含重复数字的新的整数。
// 保证输入的整数最后一位不是 0 。
// 输入描述：
// 输入一个int型整数
// 输出描述：
// 按照从右向左的阅读顺序，返回一个不含重复数字的新的整数
func 提取不重复的整数() {
	var str string
	fmt.Scanln(&str)
	strSlice := strings.Split(str, "")
	for i := 0; i < len(strSlice)/2; i++ {
		strSlice[i], strSlice[len(strSlice)-1-i] = strSlice[len(strSlice)-1-i], strSlice[i]
	}

	//map是无序的，键值对的顺序不是赋值的顺序
	strMap := map[string]string{}
	slice := []string{}
	for _, v := range strSlice {
		if _, ok := strMap[v]; !ok {
			strMap[v] = v
			slice = append(slice, v)
		}
	}
	for _, v := range slice {
		fmt.Printf("%s", strMap[v])
	}
}

// 字符个数统计
// 编写一个函数，计算字符串中含有的不同字符的个数。字符在 ASCII 码范围内( 0~127 ，包括 0 和 127 )，换行表示结束符，不算在字符里。不在范围内的不作统计。多个相同的字符只计算一次
// 例如，对于字符串 abaca 而言，有 a、b、c 三种不同的字符，因此输出 3 。
// 数据范围：
// 1≤n≤500
// 输入描述：
// 输入一行没有空格的字符串。
// 输出描述：
// 输出 输入字符串 中范围在(0~127，包括0和127)字符的种数。
func 字符个数统计() {
	var str string
	fmt.Scanln(&str)
	strSlice := strings.Split(str, "")
	strMap := map[string]string{}
	for _, v := range strSlice {
		strMap[v] = v
	}
	fmt.Println(len(strMap))
}

// 数字颠倒
// 输入一个整数，将这个整数以字符串的形式逆序输出
// 程序不考虑负数的情况，若数字含有0，则逆序形式也含有0，如输入为100，则输出为001
// 输入描述：
// 输入一个int整数
// 输出描述：
// 将这个整数以字符串的形式逆序输出
func 数字颠倒() {
	var str string
	fmt.Scanln(&str)
	strSlice := []rune(str)
	for i := 0; i < len(strSlice)/2; i++ {
		strSlice[i], strSlice[len(strSlice)-1-i] = strSlice[len(strSlice)-1-i], strSlice[i]
	}
	fmt.Println(string(strSlice))
}

// 句子逆序
// 将一个英文语句以单词为单位逆序排放。例如“I am a boy”，逆序排放后为“boy a am I”
// 所有单词之间用一个空格隔开，语句中除了英文字母外，不再包含其他字符
// 输入描述：
// 输入一个英文语句，每个单词用空格隔开。保证输入只包含空格和字母。
// 输出描述：
// 得到逆序的句子
func 句子逆序() {
	reader := bufio.NewReader(os.Stdin)
	str, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("err:%v\n", err)
		return
	}
	str = strings.TrimSuffix(str, "\n")
	strSlice := strings.Split(str, " ")
	for i := 0; i < len(strSlice)/2; i++ {
		strSlice[i], strSlice[len(strSlice)-1-i] = strSlice[len(strSlice)-1-i], strSlice[i]
	}

	for _, v := range strSlice {
		fmt.Printf("%s ", v)
	}
}

// 字符串排序
// 给定 n 个字符串，请对 n 个字符串按照字典序排列。
// 数据范围： 1≤n≤1000  ，字符串长度满足1≤len≤100
// 输入描述：
// 输入第一行为一个正整数n(1≤n≤1000),下面n行为n个字符串(字符串长度≤100),字符串中只含有大小写字母。
// 输出描述：
// 数据输出n行，输出结果为按照字典序排列的字符串。
func 字符串排序() {
	var n int
	fmt.Scanln(&n)
	var str = make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Scanln(&str[i])
	}
	sort.Strings(str)
	for _, v := range str {
		fmt.Println(v)
	}
}

// 求int型正整数在内存中存储时1的个数
// 输入一个 int 型的正整数，计算出该 int 型数据在内存中存储时 1 的个数。
// 数据范围：保证在 32 位整型数字范围内
// 输入描述：
// 输入一个整数（int类型）
// 输出描述：
// 这个数转换成2进制后，输出1的个数
func 求int型正整数在内存中存储时1的个数() {
	var n int64
	fmt.Scanln(&n)
	str := strconv.FormatInt(n, 2)
	strSlice := strings.Split(str, "")
	count := 0
	for _, v := range strSlice {
		if v == "1" {
			count++
		}
	}
	fmt.Println(count)
}

// 01背包问题
// 有N件物品和一个容量为V的背包。第i件物品的价值是C[i]，重量是W[i]，求解将哪些物品装入背包可使价值总和最大。
// 输入描述：
// 输入第一行数 N V （1<=N<=500）(1<=V<=10000)
// 输入N行两个数字，代表C W（1<=C<=50000,1<=W<=10000）。C:价值 W:重量
// 输出描述：
// 输出最大价值总和
func 背包问题() {
	//f(n,v) n:可放物品数量 v:背包容量
	//1、v < w[i]，f(n,v) = f(n-1,v)
	//2、v >= w[i]，f(n,v) = max{ f(n-1,v) , f(n-1,v-w[i])+c[i] }
	var N, V int
	fmt.Scanln(&N, &V)
	var C, W int
	var items = map[int]map[string]int{}
	for i := 1; i <= N; i++ {
		fmt.Scanln(&C, &W)
		items[i] = map[string]int{
			"C": C,
			"W": W,
		}
	}
	pd := make([][]int, N+1)
	for i := 0; i <= N; i++ {
		pd[i] = make([]int, V+1)
	}
	//i行代表有哪些物品可选
	for i := 1; i <= N; i++ {
		//j列代表不同的背包容量
		for j := 1; j <= V; j++ {
			if j < items[i]["W"] {
				pd[i][j] = pd[i-1][j]
			} else {
				pd[i][j] = max(pd[i-1][j], pd[i-1][j-items[i]["W"]]+items[i]["C"])
			}
		}
	}
	//for i := 0; i <= N; i++ {
	//	fmt.Println(pd[i])
	//}
	fmt.Println(pd[N][V])
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 购物单
type Item struct {
	value  int //价格
	weight int //满意度 = 价格*重要度
	order  int //编号
}

// todo 有BUG，后续再看
func 购物单() {
	var N, m int
	fmt.Scanln(&N, &m)
	items := make([]Item, m+1)
	var v, p, q int
	for i := 1; i <= m; i++ {
		fmt.Scanln(&v, &p, &q)

		items[i] = Item{
			value:  v,
			weight: v * p,
			order:  q,
		}
	}

	var dp = make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, N+1)
	}
	//标记对应主件是否被购买
	var mark = make([]bool, m+1)
	for i := 1; i <= m; i++ {
		for j := 1; j <= N; j++ {
			if j < items[i].value {
				dp[i][j] = dp[i-1][j]
				mark[i] = false
			} else {
				if items[i].order == 0 {
					dp[i][j] = max(dp[i-1][j], dp[i-1][j-items[i].value]+items[i].weight)
					if dp[i][j] > dp[i-1][j] {
						mark[i] = true
					} else {
						mark[i] = false
					}
				} else {
					//说明主件已被购买
					if mark[items[i].order] {
						dp[i][j] = max(dp[i-1][j], dp[i-1][j-items[i].value]+items[i].weight)
					}
					mark[i] = false
				}
			}
		}
	}

	for i := 0; i <= m; i++ {
		fmt.Println(dp[i])
	}
}

// 坐标移动
// 开发一个坐标计算工具， A表示向左移动，D表示向右移动，W表示向上移动，S表示向下移动。从（0,0）点开始移动，从输入字符串里面读取一些坐标，并将最终输入结果输出到输出文件里面。
// 输入：
// 合法坐标为A(或者D或者W或者S) + 数字（两位以内）
// 坐标之间以;分隔。
// 非法坐标点需要进行丢弃。如AA10;  A1A;  $%$;  YAD; 等。
// 输入描述：
// 一行字符串
// 输出描述：
// 最终坐标，以逗号分隔
func 坐标移动() {
	reader := bufio.NewReader(os.Stdin)
	str, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("err:%v\n", err)
		return
	}
	str = strings.TrimSuffix(str, "\n")
	strSlice := strings.Split(str, ";")
	xy := map[string]int64{
		"x": 0,
		"y": 0,
	}
	for _, v := range strSlice {
		length := len(v)
		var num int64
		if length > 3 || length <= 1 {
			continue
		}

		num, err = strconv.ParseInt(string(v[1:]), 10, 64)
		if err != nil {
			continue
		}
		switch v[0] {
		case 'A':
			xy["x"] = xy["x"] - num
		case 'D':
			xy["x"] = xy["x"] + num
		case 'W':
			xy["y"] = xy["y"] + num
		case 'S':
			xy["y"] = xy["y"] - num
		default:
			continue
		}
	}
	fmt.Printf("%d,%d", xy["x"], xy["y"])
}

// 识别有效的IP地址和掩码并进行分类统计
// 请解析IP地址和对应的掩码，进行分类识别。要求按照A/B/C/D/E类地址归类，不合法的地址和掩码单独归类。
// 所有的IP地址划分为 A,B,C,D,E五类
// A类地址从1.0.0.0到126.255.255.255;
// B类地址从128.0.0.0到191.255.255.255;
// C类地址从192.0.0.0到223.255.255.255;
// D类地址从224.0.0.0到239.255.255.255；
// E类地址从240.0.0.0到255.255.255.255
//
// 私网IP范围是：
// 从10.0.0.0到10.255.255.255
// 从172.16.0.0到172.31.255.255
// 从192.168.0.0到192.168.255.255
//
// 子网掩码为二进制下前面是连续的1，然后全是0。（例如：255.255.255.32就是一个非法的掩码）
// （注意二进制下全是1或者全是0均为非法子网掩码）
//
// 注意：
// 1. 类似于【0.*.*.*】和【127.*.*.*】的IP地址不属于上述输入的任意一类，也不属于不合法ip地址，计数时请忽略
// 2. 私有IP地址和A,B,C,D,E类地址是不冲突的
// 输入描述：
// 多行字符串。每行一个IP地址和掩码，用~隔开。
// 请参考帖子https://www.nowcoder.com/discuss/276处理循环输入的问题。
// 输出描述：
// 统计A、B、C、D、E、错误IP地址或错误掩码、私有IP的个数，之间以空格隔开。
func 识别有效的IP地址和掩码并进行分类统计() {
	scaner := bufio.NewScanner(os.Stdin)
	result := make([]int, 7)
	for scaner.Scan() {
		ips := strings.Split(scaner.Text(), "~")
		ipMask := strings.Split(ips[1], ".")
		ip := strings.Split(ips[0], ".")
		if ip[0] == "0" || ip[0] == "127" {
			continue
		}
		wrong, private := judgeIpIsWrongOrPrivate(ip, ipMask)
		if wrong {
			result[5]++
			continue
		}
		if private {
			result[6]++
		}

		ipInt, _ := strconv.ParseInt(ip[0], 10, 64)

		if 1 <= ipInt && ipInt <= 126 {
			result[0]++
		}
		if 128 <= ipInt && ipInt <= 191 {
			result[1]++
		}
		if 192 <= ipInt && ipInt <= 223 {
			result[2]++
		}
		if 224 <= ipInt && ipInt <= 239 {
			result[3]++
		}
		if 240 <= ipInt && ipInt <= 255 {
			result[4]++
		}
	}
	for _, v := range result {
		fmt.Printf("%d ", v)
	}
}

// 返回值：判断ip或子网掩码是否错误,ip是否为私有IP
func judgeIpIsWrongOrPrivate(ip, ipMask []string) (bool, bool) {
	//标记子网掩码上一段是否为全1
	var mark = false
	//标记四段是否全为0
	var all0Mark = true
	for k, v := range ipMask {
		if v != "255" && v != "254" && v != "252" && v != "248" && v != "240" && v != "224" && v != "192" && v != "128" && v != "0" {
			return true, false
		} else {
			if k != 0 {
				if !mark && v != "0" {
					return true, false
				}
				if k == 3 && v == "255" {
					return true, false
				}
			}
			if v == "255" {
				mark = true
			} else {
				mark = false
			}

			if v != "0" {
				all0Mark = false
			}
		}
	}
	if all0Mark {
		return true, false
	}

	for _, v := range ip {
		if v == "" {
			return true, false
		}
		ipInt, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return true, false
		}
		if ipInt < 0 || ipInt > 255 {
			return true, false
		}
	}
	ip1Int, _ := strconv.ParseInt(ip[0], 10, 64)
	if ip1Int == 10 {
		return false, true
	}
	ip2Int, _ := strconv.ParseInt(ip[1], 10, 64)
	if ip1Int == 172 {
		if 16 <= ip2Int && ip2Int <= 31 {
			return false, true
		}
	}
	if ip1Int == 192 && ip2Int == 168 {
		return false, true
	}
	return false, false
}

// 简单错误记录
type ErrItem struct {
	path    string
	errLine string
	count   int
}

// 开发一个简单错误记录功能小模块，能够记录出错的代码所在的文件名称和行号。
// 处理：
// 1、 记录最多8条错误记录，循环记录，最后只用输出最后出现的八条错误记录。对相同的错误记录只记录一条，但是错误计数增加。最后一个斜杠后面的带后缀名的部分（保留最后16位）和行号完全匹配的记录才做算是“相同”的错误记录。
// 2、 超过16个字符的文件名称，只记录文件的最后有效16个字符；
// 3、 输入的文件可能带路径，记录文件名称不能带路径。也就是说，哪怕不同路径下的文件，如果它们的名字的后16个字符相同，也被视为相同的错误记录
// 4、循环记录时，只以第一次出现的顺序为准，后面重复的不会更新它的出现时间，仍以第一次为准
// 输入描述：
// 每组只包含一个测试用例。一个测试用例包含一行或多行字符串。每行包括带路径文件名称，行号，以空格隔开。
// 输出描述：
// 将所有的记录统计并将结果输出，格式：文件名 代码行数 数目，一个空格隔开，如：
func 简单错误记录() {
	errSlice := []map[string]ErrItem{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		str := strings.Split(scanner.Text(), " ")
		files := strings.Split(str[0], "\\")
		fileName := files[len(files)-1]
		if len(fileName) > 16 {
			fileName = fileName[len(fileName)-16:]
		}
		mark := false
		for _, v := range errSlice {
			if item, ok := v[fileName]; ok {
				if item.errLine == str[1] {
					item.count++
					v[fileName] = item
					mark = true
					break
				}
			}
		}
		if mark {
			continue
		}
		errSlice = append(errSlice, map[string]ErrItem{
			fileName: ErrItem{
				path:    str[0],
				errLine: str[1],
				count:   1,
			},
		})
	}

	j := 0
	if len(errSlice) > 8 {
		j = len(errSlice) - 8
	}
	for i := j; i < len(errSlice); i++ {
		for name, value := range errSlice[i] {
			fmt.Printf("%s %s %d\n", name, value.errLine, value.count)
		}
	}
}

// 密码验证合格程序
// 密码要求:
// 1.长度超过8位
// 2.包括大小写字母.数字.其它符号,以上四种至少三种
// 3.不能有长度大于2的包含公共元素的子串重复 （注：其他符号不含空格或换行）
// 输入描述：
// 一组字符串。
// 输出描述：
// 如果符合要求输出：OK，否则输出NG
func 密码验证合格程序() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		str := scanner.Text()
		//要求1
		if len(str) <= 8 {
			fmt.Println("NG")
			continue
		}
		//要求2
		slice := make([]int, 4)
		for _, v := range str {
			if 'a' <= v && v <= 'z' {
				slice[0] = 1
			} else if 'A' <= v && v <= 'Z' {
				slice[1] = 1
			} else if '0' <= v && v <= '9' {
				slice[2] = 1
			} else if v != ' ' && v != '\n' {
				slice[3] = 1
			}
		}
		count := 0
		for _, v := range slice {
			if v == 1 {
				count++
			}
		}
		if count < 3 {
			fmt.Println("NG")
			continue
		}
		//要求3
		subStr := false
		for i := 0; i < len(str)-3; i++ {
			subStr = false
			for j := i + 2; j < len(str); j++ {
				if strings.Count(str, str[i:j+1]) >= 2 {
					subStr = true
					break
				}
			}
			if subStr {
				break
			}
		}
		if subStr {
			fmt.Println("NG")
			continue
		}
		fmt.Println("OK")
	}
}

// 简单密码
// 现在有一种密码变换算法。
// 九键手机键盘上的数字与字母的对应： 1--1， abc--2, def--3, ghi--4, jkl--5, mno--6, pqrs--7, tuv--8 wxyz--9, 0--0，把密码中出现的小写字母都变成九键键盘对应的数字，如：a 变成 2，x 变成 9.
// 而密码中出现的大写字母则变成小写之后往后移一位，如：X ，先变成小写，再往后移一位，变成了 y ，例外：Z 往后移是 a 。
// 数字和其它的符号都不做变换。
// 输入描述：
// 输入一组密码，长度不超过100个字符。
// 输出描述：
// 输出密码变换后的字符串
func 简单密码() {
	var str string
	fmt.Scanln(&str)
	var result = []string{}
	for _, v := range str {
		if 'A' <= v && v < 'Z' {
			result = append(result, strings.ToLower(string(v+1)))
		} else if v == 'Z' {
			result = append(result, "a")
		} else if 'a' <= v && v <= 'z' {
			if 'a' <= v && v <= 'c' {
				result = append(result, "2")
			} else if 'd' <= v && v <= 'f' {
				result = append(result, "3")
			} else if 'g' <= v && v <= 'i' {
				result = append(result, "4")
			} else if 'j' <= v && v <= 'l' {
				result = append(result, "5")
			} else if 'm' <= v && v <= 'o' {
				result = append(result, "6")
			} else if 'p' <= v && v <= 's' {
				result = append(result, "7")
			} else if 't' <= v && v <= 'v' {
				result = append(result, "8")
			} else {
				result = append(result, "9")
			}
		} else {
			result = append(result, string(v))
		}
	}
	fmt.Println(strings.Join(result, ""))
}

// 汽水瓶
// 某商店规定：三个空汽水瓶可以换一瓶汽水，允许向老板借空汽水瓶（但是必须要归还）。
// 小张手上有n个空汽水瓶，她想知道自己最多可以喝到多少瓶汽水。
// 注意：本题存在多组输入。输入的 0 表示输入结束，并不用输出结果。
// 输入描述：
// 输入文件最多包含 10 组测试数据，每个数据占一行，仅包含一个正整数 n（ 1<=n<=100 ），表示小张手上的空汽水瓶数。n=0 表示输入结束，你的程序不应当处理这一行。
// 输出描述：
// 对于每组测试数据，输出一行，表示最多可以喝的汽水瓶数。如果一瓶也喝不到，输出0。
func 汽水瓶() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		str := scanner.Text()
		if str == "0" {
			break
		}
		n, err := strconv.ParseInt(str, 10, 64)
		if err != nil {
			break
		}
		drinkNum := int64(0)
		for {
			if n < 3 {
				if n == 2 {
					n = 3
				} else {
					break
				}
			}
			currentDrinkNum := n / 3
			drinkNum = drinkNum + currentDrinkNum
			n = currentDrinkNum + n%3
		}
		fmt.Println(drinkNum)
	}
}

// 删除字符串中出现次数最少的字符
// 实现删除字符串中出现次数最少的字符，若出现次数最少的字符有多个，则把出现次数最少的字符都删除。输出删除这些单词后的字符串，字符串中其它字符保持原来的顺序。
// 输入描述：
// 字符串只包含小写英文字母, 不考虑非法输入，输入的字符串长度小于等于20个字节。
// 输出描述：
// 删除字符串中出现次数最少的字符后的字符串。
func 删除字符串中出现次数最少的字符() {
	var str string
	fmt.Scanln(&str)
	resMapStrInt := map[string]int{}
	for _, v := range str {
		if value, ok := resMapStrInt[string(v)]; ok {
			value++
			resMapStrInt[string(v)] = value
		} else {
			resMapStrInt[string(v)] = 1
		}
	}
	resMapIntStr := map[int][]string{}
	resIntSlice := []int{}
	for k, v := range resMapStrInt {
		resMapIntStr[v] = append(resMapIntStr[v], k)
		resIntSlice = append(resIntSlice, v)
	}
	sort.Ints(resIntSlice)
	strs := resMapIntStr[resIntSlice[0]]
	res := []string{}
	for _, v := range str {
		mark := false
		for _, s := range strs {
			if string(v) == s {
				mark = true
				break
			}
		}
		if mark {
			continue
		}
		res = append(res, string(v))
	}
	fmt.Println(strings.Join(res, ""))
}

// 滑动窗口
func 从字符串中查找连续最长有序子串() {
	var str string
	fmt.Scanln(&str)
	maxLength := 1
	currentLength := 1
	maxStartIndex := 0
	currentStartIndex := 0
	for i := 1; i < len(str); i++ {
		if str[i] > str[i-1] {
			currentLength++
			if currentLength > maxLength {
				maxLength = currentLength
				maxStartIndex = currentStartIndex
			}
		} else {
			currentLength = 1
			currentStartIndex = i
		}
	}
	fmt.Println(str[maxStartIndex : maxStartIndex+maxLength])
}

// 动态规划
// 1、dp[i][j]表示字符串中从0到第j个字符的子串中，第i个字符做第一个元素时 的最长的子序列。即最长的子序列在dp[i][len(str)-1]中
// 2、递推公式：
// dp[i][i] = str[i],当str[j] > dp[i][j-1][len(dp[i][j-1])-1]时，dp[i][j] = dp[i][j-1] + str[j]；否则，dp[i][j] = dp[i][j-1]
func 从字符串中查找非连续最长有序子序列() {
	var str string
	fmt.Scanln(&str)
	n := len(str)
	dp := make([][]string, n+1)
	for k, v := range dp {
		v = make([]string, n+1)
		dp[k] = v
	}
	for i := 1; i <= n; i++ {
		dp[i][i] = string(str[i-1])
		for j := i + 1; j <= n; j++ {
			if str[j-1] > dp[i][j-1][len(dp[i][j-1])-1] {
				dp[i][j] = dp[i][j-1] + string(str[j-1])
			} else {
				dp[i][j] = dp[i][j-1]
			}
		}
	}

	maxLength := 1
	resultStr := ""
	for i := 1; i <= n; i++ {
		if len(dp[i][n]) > maxLength {
			maxLength = len(dp[i][n])
			resultStr = dp[i][n]
		}
	}
	fmt.Println(maxLength)
	fmt.Println(resultStr)
}

// 动态规划
// 动态规划的题一般就两个要点：1、确定格子怎么画。2、找到递推公示
// 1、本题格子i,j分别表示第一个字符串的游标和第二个字符串的游标，dp[i][j]表示第一个字符串从0到第i个字符的子串与第二个字符串从0到第j个字符串的子串的公共子串最大的长度。
// 2、递推公式为：当str1[i] == str2[j]时，dp[i][j] = dp[i-1][j-1]+1；否则，dp[i][j] = max(dp[i-1][j],dp[i][j-1])
func 从两个字符串中找出公共最长子序列() {
	var str1, str2 string
	fmt.Scanln(&str1)
	fmt.Scanln(&str2)
	n := len(str1)
	m := len(str2)
	dp := make([][]int, n+1)
	dpStr := make([][]string, n+1)
	for k, v := range dp {
		v = make([]int, m+1)
		dp[k] = v
	}
	for k, v := range dpStr {
		v = make([]string, m+1)
		dpStr[k] = v
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if str1[i-1] == str2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
				dpStr[i][j] = dpStr[i-1][j-1] + string(str1[i-1])
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
				if dp[i][j] == dp[i-1][j] {
					dpStr[i][j] = dpStr[i-1][j]
				} else {
					dpStr[i][j] = dpStr[i][j-1]
				}
			}
		}
	}
	fmt.Println(dp[n][m])
	fmt.Println(dpStr[n][m])
}

// 合唱队
func 合唱队() {
	//f(n-1) < f(n) > f(n+1)
	var N int
	fmt.Scanln(&N)
	reader := bufio.NewReader(os.Stdin)
	str, _ := reader.ReadString('\n')
	str = strings.TrimSuffix(str, "\n")
	strSlice := strings.Split(str, " ")
	pd := make([]int, N+1)
	for i := 2; i < N; i++ {
		strSlice1 := strSlice[:i]
		for j := 0; j < len(strSlice1); j++ {
			strSlice1[j], strSlice1[len(strSlice1)-1-j] = strSlice1[len(strSlice1)-1-j], strSlice1[j]
		}
		strSlice2 := strSlice[i-1:]
		pd[i] = strSliceMaxLength(strSlice1) + strSliceMaxLength(strSlice2) - 1
	}

	maxCount := 0
	for _, v := range pd {
		if v > maxCount {
			maxCount = v
		}
	}
	fmt.Println(N - maxCount)
}

// 获取字符串切片中第一个元素作为起始元素的最长递减子序列
// 动态规划
// pd[i]表示从0到i个元素组成的字符串切片中以0个元素为起点的递减子序列最大元素个数
// 递推公式：dp[1] = 1,当str[i] < str[i-1]时，dp[i] = dp[i-1]+1，否则，pd[i] = dp[i-1]
func strSliceMaxLength(strSlice []string) int {
	dp := make([]int, len(strSlice)+1)
	dp[1] = 1
	for i := 1; i < len(strSlice); i++ {
		intCurrent, _ := strconv.Atoi(strSlice[i])
		intLast, _ := strconv.Atoi(strSlice[i-1])
		if intCurrent < intLast {
			dp[i+1] = dp[i] + 1
		} else {
			dp[i+1] = dp[i]
		}
	}
	return dp[len(dp)-1]
}

// HJ26字符串排序
// 编写一个程序，将输入字符串中的字符按如下规则排序。
// 规则 1 ：英文字母从 A 到 Z 排列，不区分大小写。
// 如，输入： Type 输出： epTy
// 规则 2 ：同一个英文字母的大小写同时存在时，按照输入顺序排列。
// 如，输入： BabA 输出： aABb
// 规则 3 ：非英文字母的其它字符保持原来的位置。
// 如，输入： By?e 输出： Be?y
// 输入描述：
// 输入字符串
// 输出描述：
// 输出字符串
func HJ26字符串排序() {
	reader := bufio.NewReader(os.Stdin)
	str, _ := reader.ReadString('\n')
	str = strings.TrimSuffix(str, "\n")
	mark := map[int]string{}
	strSlice := map[string][]string{}
	for k, v := range str {
		if ('a' <= v && v <= 'z') || ('A' <= v && v <= 'Z') {
			strSlice[strings.ToLower(string(v))] = append(strSlice[strings.ToLower(string(v))], string(v))
		} else {
			mark[k] = string(v)
		}
	}

	markStr := []string{}
	for k, _ := range strSlice {
		markStr = append(markStr, k)
	}
	sort.Strings(markStr)

	result := ""
	markI := 0
	for _, v := range markStr {
		for _, value := range strSlice[v] {
			for {
				if v, ok := mark[markI]; ok {
					result += v
					markI++
				} else {
					break
				}
			}
			result += value
			markI++
		}
	}
	for i := len(result); i < len(str); i++ {
		if v, ok := mark[i]; ok {
			result += v
		}
	}

	fmt.Println(result)
}

// HJ27查找兄弟单词
// 定义一个单词的“兄弟单词”为：交换该单词字母顺序（注：可以交换任意次），而不添加、删除、修改原有的字母就能生成的单词。
// 兄弟单词要求和原来的单词不同。例如： ab 和 ba 是兄弟单词。 ab 和 ab 则不是兄弟单词。
// 现在给定你 n 个单词，另外再给你一个单词 x ，让你寻找 x 的兄弟单词里，按字典序排列后的第 k 个单词是什么？
// 注意：字典中可能有重复单词。
// 输入描述：
// 输入只有一行。 先输入字典中单词的个数n，再输入n个单词作为字典单词。 然后输入一个单词x 最后后输入一个整数k
// 输出描述：
// 第一行输出查找到x的兄弟单词的个数m 第二行输出查找到的按照字典顺序排序后的第k个兄弟单词，没有符合第k个的话则不用输出。
func HJ27查找兄弟单词() {
	reader := bufio.NewReader(os.Stdin)
	str, _ := reader.ReadString('\n')
	str = strings.TrimSuffix(str, "\n")
	strSlice := strings.Split(str, " ")
	brotherStr := strSlice[len(strSlice)-2]
	n, _ := strconv.Atoi(strSlice[len(strSlice)-1])
	strSlice = strSlice[1 : len(strSlice)-2]
	brotherStrMap := map[string]int{}
	for _, v := range brotherStr {
		if _, ok := brotherStrMap[string(v)]; ok {
			brotherStrMap[string(v)]++
		} else {
			brotherStrMap[string(v)] = 1
		}
	}

	brotherSlice := []string{}
	for _, v := range strSlice {
		if len(v) != len(brotherStr) {
			continue
		}
		if v == brotherStr {
			continue
		}
		currentBrotherMap := map[string]int{}
		for _, i := range v {
			if _, ok := currentBrotherMap[string(i)]; ok {
				currentBrotherMap[string(i)]++
			} else {
				currentBrotherMap[string(i)] = 1
			}
		}
		isBrother := true
		for k, v := range currentBrotherMap {
			if _, ok := brotherStrMap[k]; !ok {
				isBrother = false
				break
			}
			if brotherStrMap[k] != v {
				isBrother = false
				break
			}
		}
		if !isBrother {
			continue
		}
		brotherSlice = append(brotherSlice, v)
	}
	sort.Strings(brotherSlice)
	brotherNum := len(brotherSlice)
	fmt.Println(brotherNum)
	if n < brotherNum {
		fmt.Println(brotherSlice[n-1])
	}
}

// HJ29字符串加解密
// 对输入的字符串进行加解密，并输出。
// 加密方法为：
// 当内容是英文字母时则用该英文字母的后一个字母替换，同时字母变换大小写,如字母a时则替换为B；字母Z时则替换为a；
// 当内容是数字时则把该数字加1，如0替换1，1替换2，9替换0；
// 其他字符不做变化。
// 解密方法为加密的逆过程。
// 输入描述：
// 第一行输入一串要加密的密码
// 第二行输入一串加过密的密码
// 输出描述：
// 第一行输出加密后的字符
// 第二行输出解密后的字符
func HJ29字符串加解密() {
	scanner := bufio.NewScanner(os.Stdin)
	var mode rune = 1
	for scanner.Scan() {
		fmt.Println(envStr(scanner.Text(), mode))
		mode = -1
	}
}
func envStr(str string, mode rune) string {
	result := ""
	//mod:1，表示要加密;mod:-1,表示解密
	for _, v := range str {
		if 'a' <= v && v <= 'z' {
			if mode == 1 {
				if v == 'z' {
					result += "A"
					continue
				}
			} else {
				if v == 'a' {
					result += "Z"
					continue
				}
			}
			result += strings.ToUpper(string(v + mode))
			continue
		}
		if 'A' <= v && v <= 'Z' {
			if mode == 1 {
				if v == 'Z' {
					result += "a"
					continue
				}
			} else {
				if v == 'A' {
					result += "z"
					continue
				}
			}
			result += strings.ToLower(string(v + mode))
			continue
		}
		intS, _ := strconv.Atoi(string(v))
		if 0 <= intS && intS <= 9 {
			if mode == 1 {
				if intS == 9 {
					result += "0"
					continue
				}
			} else {
				if intS == 0 {
					result += "9"
					continue
				}
			}
			result += strconv.FormatInt(int64(intS+int(mode)), 10)
			continue
		}
	}
	return result
}

// HJ30字符串合并处理
// 按照指定规则对输入的字符串进行处理。
// 详细描述：
// 第一步：将输入的两个字符串str1和str2进行前后合并。如给定字符串 "dec" 和字符串 "fab" ， 合并后生成的字符串为 "decfab"
// 第二步：对合并后的字符串进行排序，要求为：下标为奇数的字符和下标为偶数的字符分别从小到大排序。这里的下标的意思是字符在字符串中的位置。注意排序后在新串中仍需要保持原来的奇偶性。例如刚刚得到的字符串“decfab”，分别对下标为偶数的字符'd'、'c'、'a'和下标为奇数的字符'e'、'f'、'b'进行排序（生成 'a'、'c'、'd' 和 'b' 、'e' 、'f'），再依次分别放回原串中的偶数位和奇数位，新字符串变为“abcedf”
// 第三步：对排序后的字符串中的'0'~'9'、'A'~'F'和'a'~'f'字符，需要进行转换操作。
// 转换规则如下：
// 对以上需要进行转换的字符所代表的十六进制用二进制表示并倒序，然后再转换成对应的十六进制大写字符（注：字符 a~f 的十六进制对应十进制的10~15，大写同理）。
// 如字符 '4'，其二进制为 0100 ，则翻转后为 0010 ，也就是 2 。转换后的字符为 '2'。
// 如字符 ‘7’，其二进制为 0111 ，则翻转后为 1110 ，对应的十进制是14，转换为十六进制的大写字母为 'E'。
// 如字符 'C'，代表的十进制是 12 ，其二进制为 1100 ，则翻转后为 0011，也就是3。转换后的字符是 '3'。
// 根据这个转换规则，由第二步生成的字符串 “abcedf” 转换后会生成字符串 "5D37BF"。
// 输入描述：
// 样例输入两个字符串，用空格隔开。
// 输出描述：
// 输出转化后的结果
func HJ30字符串合并处理() {
	var a, b string
	fmt.Scan(&a, &b)
	//第一步：合并
	str := a + b
	//第二步：排序
	var i, j = []string{}, []string{}
	for k, v := range str {
		if (k+1)%2 == 1 {
			i = append(i, string(v))
		} else {
			j = append(j, string(v))
		}
	}
	sort.Strings(i)
	sort.Strings(j)
	str = ""
	for z := 0; z < len(i)+len(j); z++ {
		if z < len(i) {
			str += i[z]
		}
		if z < len(j) {
			str += j[z]
		}
	}
	//第三步：转换
	result := ""
	for _, v := range str {
		if !(('0' <= v && v <= '9') || ('a' <= v && v <= 'f') || ('A' <= v && v <= 'F')) {
			result += string(v)
			continue
		}
		intT, _ := strconv.ParseInt(string(v), 16, 64)
		twoString := strconv.FormatInt(intT, 2)
		twoStringStr := ""
		for i := len(twoString); i < 4; i++ {
			twoStringStr += "0"
		}
		twoStringStr += twoString
		twoStringSlice := []rune(twoStringStr)
		for i := 0; i < len(twoStringSlice)/2; i++ {
			twoStringSlice[i], twoStringSlice[len(twoStringSlice)-1-i] = twoStringSlice[len(twoStringSlice)-1-i], twoStringSlice[i]
		}
		twoStringStr = string(twoStringSlice)

		intTen, _ := strconv.ParseInt(twoStringStr, 2, 64)
		result += strings.ToUpper(strconv.FormatInt(intTen, 16))
	}
	fmt.Println(result)
}

// 宝宝和妈妈参加亲子游戏，在一个二维矩阵（N*N）的格子地图上，宝宝和妈妈抽签决定各自的位置，地图上每个格子有不同的糖果数量，部分格子有障碍物。
// 游戏规则是妈妈必须在最短的时间（每个单位时间只能走一步）到达宝宝的位置，路上的所有糖果都可以拿走，不能走障碍物的格子，只能上下左右走。
// 请问妈妈在最短到达宝宝位置的时间内最多拿到多少糖果（优先考虑最短时间到达的情况下尽可能多拿糖果）。
// 输入描述：
// 第一行输入为 N，N 表示二维矩阵的大小。 之后 N 行，每行有 N 个值，表格矩阵每个位置的值，其中:
// -3: 妈妈
// -2: 宝宝
// -1: 障碍
// ≥0 : 糖果数(0 表示没有糖果，但是可以走)
// 输出描述：
// 输出妈妈在最短到达宝宝位置的时间内最多拿到多少糖果，行未无多余空格。
// 广度搜索(BFS)
type Point struct {
	x, y, candies int
}

func 亲子游戏() {
	var N int
	fmt.Scanln(&N)
	matrix := make([][]int, N)
	var startX, startY, targetX, targetY int
	for i := 0; i < N; i++ {
		matrix[i] = make([]int, N)
		for j := 0; j < N; j++ {
			fmt.Scan(&matrix[i][j])
			if matrix[i][j] == -3 {
				startX = i
				startY = j
			} else if matrix[i][j] == -2 {
				targetX = i
				targetY = j
			}
		}
	}
	fmt.Println(bfs亲子游戏(matrix, startX, startY, targetX, targetY))

	//maxCandies := -1
	//visit := make([][]bool, N)
	//for i := 0; i < N; i++ {
	//	visit[i] = make([]bool, N)
	//}
	//dfs亲子游戏(matrix, startX, startY, targetX, targetY, N, 0, visit, &maxCandies)
	//fmt.Println(maxCandies)
}
func bfs亲子游戏(matrix [][]int, startX, startY, targetX, targetY int) int {
	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	//directions := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	N := len(matrix)
	visit := make([][]bool, N)
	for i := 0; i < N; i++ {
		visit[i] = make([]bool, N)
	}
	maxCandies := -1
	q := list.New()
	q.PushBack(Point{startX, startY, 0})
	for q.Len() > 0 {
		point := q.Remove(q.Front()).(Point)
		if point.x == targetX && point.y == targetY {
			maxCandies = max(maxCandies, point.candies)
			continue
		}

		for _, d := range directions {
			x := point.x + d[0]
			y := point.y + d[1]
			if x >= 0 && x < N && y >= 0 && y < N && matrix[x][y] != -1 && visit[x][y] == false {
				visit[x][y] = true
				q.PushBack(Point{x, y, point.candies + max(0, matrix[x][y])})
			}
		}
	}
	return maxCandies
}

//func dfs亲子游戏(matrix [][]int, startX, startY, targetX, targetY, N, candies int, visit [][]bool, maxCandies *int) {
//	if startX < 0 || startX >= N || startY < 0 || startY >= N || matrix[startX][startY] == -1 || visit[startX][startY] {
//		return
//	}
//
//	if startX == targetX && startY == targetY {
//		*maxCandies = max(*maxCandies, candies)
//		return
//	}
//	visit[startX][startY] = true
//	candies += max(0, matrix[startX][startY])
//
//	dfs亲子游戏(matrix, startX-1, startY, targetX, targetY, N, candies, visit, maxCandies)
//	dfs亲子游戏(matrix, startX+1, startY, targetX, targetY, N, candies, visit, maxCandies)
//	dfs亲子游戏(matrix, startX, startY-1, targetX, targetY, N, candies, visit, maxCandies)
//	dfs亲子游戏(matrix, startX, startY+1, targetX, targetY, N, candies, visit, maxCandies)
//
//	return
//}

// 给定一个字符串，只包含字母和数字，按要求找出字符串中的最长（连续）子串的长度，字符串本身是其最长的子串，
// 子串要求：
// 1、 只包含1个字母(a~z, A~Z)，其余必须是数字；
// 2、 字母可以在子串中的任意位置；如果找不到满足要求的子串，如全是字母或全是数字，则返回-1。
// 滑动窗口
func 最长连续子串() {
	var str string
	fmt.Scanln(&str)
	currentLength := 0
	maxLength := 0
	startIndex := 0
	maxStartIndex := 0
	mark := false
	for k, v := range str {
		if !mark {
			if ('a' <= v && v <= 'z') || ('A' <= v && v <= 'Z') {
				mark = true
			}
			currentLength++
			if currentLength > maxLength {
				maxStartIndex = startIndex
				maxLength = currentLength
			}
		} else {
			if ('a' <= v && v <= 'z') || ('A' <= v && v <= 'Z') {
				startIndex = k
				currentLength = 1
			} else {
				currentLength++
				if currentLength > maxLength {
					maxStartIndex = startIndex
					maxLength = currentLength
				}
			}
		}
	}
	//全为字母或全为数字
	if !mark || (mark && maxLength == 1) {
		fmt.Println(-1)
		return
	}
	fmt.Println(maxLength)
	fmt.Println(str[maxStartIndex : maxStartIndex+maxLength])
}

// 疫情期间，需要大家保证一定的社交距离，公司组织开交流会议，座位有一排共N个座位，编号分别为[0..N-1]，要求员工一个接着一个进入会议室，并且可以在任何时候离开会议室。
// 满足：每当一个员工进入时，需要坐到最大社交距离的座位（例如：位置A与左右有员工落座的位置距离分别为2和2，位置B与左右有员工落座的位置距离分别为2和3，影响因素都为2个位置，则认为座位A和B与左右位置的社交距离是一样的）；如果有多个这样的座位，则坐到索引最小的那个座位。
// 输入描述：
// 会议室座位总数 seatNum 。(1 <= seatNum <= 500)
// 员工的进出顺序 seatOrLeave 数组，元素值为 1，表示进场；元素值为负数，表示出场（特殊：位置 0 的员工不会离开）。
// 例如 - 4 表示坐在位置 4 的员工离开（保证有员工坐在该座位上）
// 输出描述：
// 最后进来员工，他会坐在第几个位置，如果位置已满，则输出 - 1 。
// 哈希表、树
func 最大社交距离() {
	var seatNum int
	fmt.Scanln(&seatNum)
	reader := bufio.NewReader(os.Stdin)
	seatOrLeave, _ := reader.ReadString('\n')
	seatOrLeave = strings.TrimSuffix(seatOrLeave, "\n")
	seatOrLeaveSlice := strings.Split(seatOrLeave, " ")
	seatOrLeaveSliceInt := []int{}
	for _, v := range seatOrLeaveSlice {
		vInt, _ := strconv.ParseInt(v, 10, 64)
		seatOrLeaveSliceInt = append(seatOrLeaveSliceInt, int(vInt))
	}
	fmt.Println(findSeat(seatNum, seatOrLeaveSliceInt))
}
func findSeat(seatNum int, seatOrLeave []int) int {
	seats := []int{}
	for _, action := range seatOrLeave {
		maxDisk := -1
		seat := -1
		pre := 0
		if action > 0 {
			//进场
			if len(seats) == 0 {
				seats = append(seats, 0)
			} else {
				for i := 0; i < seatNum; i++ {
					if !slices.Contains(seats, i) {
						left := i - pre
						right := 500
						for j := i + 1; j < seatNum; j++ {
							if slices.Contains(seats, j) {
								right = j - i
								break
							}
						}
						disk := min(left, right)
						if disk > maxDisk {
							maxDisk = disk
							seat = i
						}
					} else {
						pre = i
					}
				}
				if maxDisk == -1 {
					return -1
				}
				seats = append(seats, seat)
			}
		} else {
			//出场
			leaveSeat := -action
			for i, v := range seats {
				if v == leaveSeat {
					seats = append(seats[:i], seats[i+1:]...)
					break
				}
			}
		}
	}

	return seats[len(seats)-1]
}

// 给定两个只包含数字的数组a, b, 调整数组a里面数字的顺序，使得尽可能多的a[i] > b[i]。数组a和b中的数字各不相同。
// 输出所有可以达到最优结果的a数组数量
// 不会 todo
func 调整数组() {
	scanner := bufio.NewScanner(os.Stdin)
	str := []string{}
	for scanner.Scan() {
		str = append(str, scanner.Text())
	}
	aS := strings.Split(str[0], " ")
	bS := strings.Split(str[1], " ")
	a := []int{}
	b := []int{}
	for _, v := range aS {
		vI, _ := strconv.ParseInt(v, 10, 64)
		a = append(a, int(vI))
	}
	for _, v := range bS {
		vI, _ := strconv.ParseInt(v, 10, 64)
		b = append(b, int(vI))
	}
	fmt.Println(countOptimalArrangements(a, b))
}
func countOptimalArrangements(a, b []int) int {
	sort.Ints(a)
	sort.Ints(b)
	count := 0
	permutations := generatePermutations(a)

	for _, perm := range permutations {
		if isOptimal(perm, b) {
			count++
		}
	}

	return count
}
func generatePermutations(arr []int) [][]int {
	var res [][]int
	var permute func([]int, int)
	permute = func(arr []int, l int) {
		if l == len(arr) {
			temp := make([]int, len(arr))
			copy(temp, arr)
			res = append(res, temp)
		} else {
			for i := l; i < len(arr); i++ {
				arr[l], arr[i] = arr[i], arr[l]
				permute(arr, l+1)
				arr[l], arr[i] = arr[i], arr[l]
			}
		}
	}
	permute(arr, 0)
	return res
}
func isOptimal(a, b []int) bool {
	for i := 0; i < len(a); i++ {
		if a[i] <= b[i] {
			return false
		}
	}
	return true
}

// 攀登者喜欢寻找各种地图，并且尝试攀登到最高的山峰。
// 地图表示为一维数组，数组的索引代表水平位置，数组的元素代表相对海拔高度。其中数组元素0代表地面。
// 例如：[0,1,2,4,3,1,0,0,1,2,3,1,2,1,0]，代表如下图所示的地图，地图中有两个山脉位置分别为 1,2,3,4,5 和 8,9,10,11,12,13，最高峰高度分别为 4,3。最高峰位置分别为3,10。
// 一个山脉可能有多座山峰(高度大于相邻位置的高度，或在地图边界且高度大于相邻的高度)。登山者想要知道一张地图中有多少座山
// 输入描述
// 输入为一个整型数组，数组长度大于1。
// 输出描述
// 输出地图中山峰的数量。
func 山脉的个数() {
	reader := bufio.NewReader(os.Stdin)
	str, _ := reader.ReadString('\n')
	str = strings.TrimSuffix(str, "\n")
	strSlice := strings.Split(str, " ")
	strSliceInt := []int{}
	for _, v := range strSlice {
		vI, _ := strconv.ParseInt(v, 10, 64)
		strSliceInt = append(strSliceInt, int(vI))
	}

	count := 0
	for i := 0; i < len(strSliceInt); i++ {
		if i-1 >= 0 && i+1 < len(strSliceInt) && strSliceInt[i] > strSliceInt[i-1] && strSliceInt[i] > strSliceInt[i+1] {
			count++
			continue
		}
		if i == 0 && i+1 < len(strSliceInt) && strSliceInt[i] > strSliceInt[i+1] {
			count++
			continue
		}
		if i == len(strSliceInt)-1 && i-1 >= 0 && strSliceInt[i] > strSliceInt[i-1] {
			count++
			continue
		}
	}
	fmt.Println(count)
}

// 中秋节公司分月饼，m个员工，买了n个月饼，m<=n，每个员工至少分1个月饼，但可以分多个，单人分到最多月饼的个数为Max1，单人分到第二多月饼的个数是Max2，Max1-Max2<=3。
// 单人分到第n-1多月饼的个数是Max(n-1)，单人分到第n多月饼的个数是Max(n)，Max(n-1)-Max(n)<=3。问有多少种分月饼的方法？
// 输入描述：
// 每一行输入m n,表示m个员工，n个月饼，m<=n
// 输出描述：
// 输出有多少种月饼分法
// 示例：
// 输入
// 2 4
// 输出
// 2
// 动态规划 dp[m][n] = dp[m][n]+dp[m-1][n-k]
func 中秋节分月饼() {
	var m, n int
	fmt.Scan(&m, &n)
	fmt.Println(countWays(m, n, 1))
}

// 表示m个人分n个月饼，最少每人i个
func countWays(m, n, i int) int {
	if m == 1 {
		if n < i || n > i+3 {
			return 0
		}
		return 1
	}
	count := 0
	for j := i; j <= n/m && j <= i+3; j++ {
		count += countWays(m-1, n-j, j)
	}
	return count
}

//	  6
//	 / \
//	7   9
//	 \  /
//	 -2 6
//
// 请由该二叉树生成一个新的二叉树，它满足其树中的每个节点将包含原始树中的左子树和右子树的和。
//
//		  20 (7-2+9+6)
//	    /   \
//	   -2    6
//	    \   /
//	    0  0
//
// 左子树表示该节点左侧叶子节点为根节点的一颗新树；右子树表示该节点右侧叶子节点为根节点的一颗新树
// 输入描述：
// 2行整数，第1行表示二叉树的中序遍历，第2行表示二叉树的前序遍历，以空格分割
// 输出描述：
// 1行整数，表示求和树的中序遍历，以空格分割
type node struct {
	value int
	left  *node
	right *node
}

// 利用中序、前序特点构建二叉树
func treeBuild(inorder []int, preorder []int) *node {
	if len(preorder) == 0 {
		return nil
	}
	//前序的第一个元素就是根节点
	root := &node{value: preorder[0]}

	rootIndex := 0
	for i, v := range inorder {
		if v == preorder[0] {
			rootIndex = i
			break
		}
	}
	root.left = treeBuild(inorder[:rootIndex], preorder[1:1+rootIndex])
	root.right = treeBuild(inorder[rootIndex+1:], preorder[1+rootIndex:])
	return root
}

// 左->根->右
func inOrder(node *node) {
	if node == nil {
		return
	}
	inOrder(node.left)
	fmt.Printf("%d ", node.value)
	inOrder(node.right)
}

// 根->左->右
func preOrder(node *node) {
	if node == nil {
		return
	}
	fmt.Printf("%d ", node.value)
	preOrder(node.left)
	preOrder(node.right)
}

// 重新计算二叉树各节点值并更新
func sumSubTree(root *node) int {
	if root == nil {
		return 0
	}
	leftSum := sumSubTree(root.left)
	rightSum := sumSubTree(root.right)
	originValue := root.value
	root.value = leftSum + rightSum
	return originValue + root.value
}
func 计算二叉树() {
	scanner := bufio.NewScanner(os.Stdin)
	str := []string{}
	i := 0
	for scanner.Scan() {
		str = append(str, scanner.Text())
		i++
		if i == 2 {
			break
		}
	}
	inorder := []int{}
	preorder := []int{}
	inslice := strings.Split(str[0], " ")
	preslice := strings.Split(str[1], " ")
	for _, v := range inslice {
		in, _ := strconv.ParseInt(v, 10, 64)
		inorder = append(inorder, int(in))
	}
	for _, v := range preslice {
		pre, _ := strconv.ParseInt(v, 10, 64)
		preorder = append(preorder, int(pre))
	}

	root := treeBuild(inorder, preorder)
	sumSubTree(root)
	inOrder(root)
}

// 给你一个字符串s,首尾相连成一个环形,请你在环中找出包含字符'l'、'o'、'x'恰好出现偶数次的最长子串的长度
// 输入描述
// 输入由一个小写字母组成的字符串s
// 输出描述
// 输出是一个整数
func 环中最长子串() {
	var s string
	fmt.Scanln(&s)
	length := len(s)
	mark := map[int]int{0: -1}
	maxLength := 0
	status := 0
	ss := s + s
	for k, v := range ss {
		switch v {
		case 'l':
			status ^= 1
		case 'o':
			status ^= 2
		case 'x':
			status ^= 4
		}

		if markValue, ok := mark[status]; ok {
			if k-markValue > maxLength && k-markValue <= length {
				maxLength = k - markValue
			}
		} else {
			mark[status] = k
		}
	}
	fmt.Println(maxLength)
}

// 一只贪吃的猴子，来到一个果园，发现许多串香蕉排成一行，每串香蕉上有若干根香蕉。每串香蕉的根数由数组numbers给出。猴子获取香蕉，每次都只能从行的开头或者末尾获取，并且只能获取N次，求猴子最多能获取多少根香蕉。
// 输入描述
// 第一行为数组numbers的长度
// 第二行为数组numbers的值每个数字通过空格分开
// 第三行输入为N，表示获取的次数
// 输出描述
// 按照题目要求能获取的最大数值
// 贪心算法
func 贪吃的猴子() {
	var length int
	var str string
	var N int
	fmt.Scan(&length)
	reader := bufio.NewReader(os.Stdin)
	str, _ = reader.ReadString('\n')
	str = strings.TrimSuffix(str, "\n")
	strSlice := strings.Split(str, " ")
	fmt.Scanln(&N)
	endIndex := length - 1
	var count int64 = 0
	for i := 0; i < length; i++ {
		startValue, _ := strconv.ParseInt(strSlice[i], 10, 64)
		endValue, _ := strconv.ParseInt(strSlice[endIndex], 10, 64)
		if startValue > endValue {
			count += startValue
		} else {
			count += endValue
			endIndex--
			i--
		}
		N--
		if N == 0 || endIndex <= i {
			break
		}
	}
	fmt.Println(count)
}

// 小扇和小船今天又玩起来了数字游戏，小船给小扇一个正整数n (1<=n<=1e9)，小扇需要找到一个比n大的数字m，使得m和n对应的二进制中1的个数要相同（如4对应二进制100,8对应二进制1000,1的个数都为1），现在求m的最小值。
func 小扇和小船的数字游戏() {
	var n int64
	fmt.Scan(&n)
	nStr := strconv.FormatInt(n, 2)
	nStr = "1" + nStr
	markIndex := 0
	for i := len(nStr) - 1; i >= 0; i-- {
		if nStr[i] == '1' {
			markIndex = i
			break
		}
	}
	mStr := nStr[:markIndex] + "0" + nStr[markIndex+1:]
	m, _ := strconv.ParseInt(mStr, 2, 64)
	fmt.Println(m)
}

// 假定街道是棋盘型的，每格距离相等，车辆通过每格街道需要时间均为 timePerRoad；
// 街道的街口（交叉点）有交通灯，灯的周期 T(=lights[row][col]) 各不相同；
// 车辆可直行、左转和右转，其中直行和左转需要等相应 T 时间的交通灯才可通行，右转无需等待。
// 现给出 n*m 个街口的交通灯周期，以及起止街口的坐标，计算车辆经过两个街口的最短时间。
// 其中：
// 1）起点和终点的交通灯不计入时间，且可以任意方向经过街口
// 2）不可超出 n*m 个街口，不可跳跃，但边线也是道路（即 lights[0][0] -> lights[0][1] 是有效路径）
// 入口函数定义：
// /**
// * lights : n*m 个街口每个交通灯的周期，值范围[0,120]，n和m的范围为[1,9]
// * timePerRoad : 相邻两个街口之间街道的通过时间,范围为[0,600]
// * rowStart : 起点的行号
// * colStart : 起点的列号
// * rowEnd : 终点的行号
// * colEnd :
// 输入描述：
// 第一行输入n和m，以空格分割
// 之后n行输入lights矩阵，矩阵每行m个整数
func 路口最短时间问题() {

}

// 实现一个模拟目录管理功能的软件，输入一个命令序列，输出最后一条命令运行结果。
// 支持命令：
// 1）创建目录命令：mkdir 目录名称，如mkdir abc为在当前目录创建abc目录，如果已存在同名目录则不执行任何操作。此命令无输出。
// 2）进入目录命令：cd 目录名称, 如cd abc为进入abc目录，特别地，cd ..为返回上级目录，如果目录不存在则不执行任何操作。此命令无输出。
// 3）查看当前所在路径命令：pwd，输出当前路径字符串。
// 约束：
// 1）目录名称仅支持小写字母；mkdir和cd命令的参数仅支持单个目录，如：mkdir abc和cd abc；不支持嵌套路径和绝对路径，如mkdir abc/efg, cd abc/efg, mkdir /abc/efg, cd /abc/efg是不支持的。
// 2）目录符号为/，根目录/作为初始目录。
// 3）任何不符合上述定义的无效命令不做任何处理并且无输出
// 输入描述：
// 输入N行字符串，每一行字符串是一条命令。
// 输出描述：
// 输出最后一条命令运行结果字符串。
type dirNode struct {
	name     string
	children map[string]*dirNode
	path     string
}

func 模拟目录管理() {
	node := &dirNode{
		name:     "/",
		children: map[string]*dirNode{},
		path:     "/",
	}
	preNode := node
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		str := scanner.Text()
		strSlice := strings.Split(str, " ")
		switch strSlice[0] {
		case "mkdir":
			if _, ok := node.children[strSlice[1]]; !ok {
				node.children = map[string]*dirNode{
					strSlice[1]: &dirNode{
						name:     strSlice[1],
						children: map[string]*dirNode{},
						path:     node.path + strSlice[1] + "/",
					},
				}
			}
		case "cd":
			if child, ok := node.children[strSlice[1]]; ok {
				preNode = node
				node = child
			} else {
				if strSlice[1] == ".." {
					node = preNode
				}
			}
		case "pwd":
			fmt.Println(node.path)
		}
	}
}

// 某部门计划通过结队编程来进行项目开发。
// 已知该部门有N名员工，每个员工有独一无二的职级，每三个员工形成一个小组进行结队编程，结队分组规则如下：
// 从部门中选出序号分别为i、j、k的三名员工，他们的职级分别为level[i]，level[j]，level[k]，结队小组满足level[i]<level[j]<level[k]或者level[i]>level[j]>level[k]，其中0<=i<j<k<n。
// 请你按上述条件计算可能组合的小组数量。同一员工可以参加多个小组。
// 输入描述：
// 第一行输入：员工总数n
// 第二行输入：按序号依次排列的员工职级level，中间用空格隔开
// 限制：1<=n<=6000，1<=level[i]<=10^5
// 输出描述：
// 可能结队的小组数量
func 结队编程() {
	var n int
	fmt.Scan(&n)
	reader := bufio.NewReader(os.Stdin)
	str, _ := reader.ReadString('\n')
	str = strings.TrimSuffix(str, "\n")
	strSlice := strings.Split(str, " ")
	count := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for z := j + 1; z < n; z++ {
				iInt, _ := strconv.ParseInt(strSlice[i], 10, 64)
				jInt, _ := strconv.ParseInt(strSlice[j], 10, 64)
				zInt, _ := strconv.ParseInt(strSlice[z], 10, 64)
				if (iInt < jInt && jInt < zInt) || (iInt > jInt && jInt > zInt) {
					count++
				}
			}
		}
	}
	fmt.Println(count)
}

// 在一个机房中，服务器的位置标识在 n*m 的整数矩阵网格中，1 表示单元格上有服务器，0 表示没有。如果两台服务器位于同一行或者同一列中紧邻的位置，则认为它们之间可以组成一个局域网。
// 请你统计机房中最大的局域网包含的服务器个数。
// 输入描述
// 第一行输入两个正整数，n和m，0<n,m<=100
// 之后为n*m的二维数组，代表服务器信息
// 输出描述
// 最大局域网包含的服务器个数。
func 可以组成网络的服务器() {
	var n, m int
	fmt.Scan(&n, &m)
	server := make([][]int, n)
	for i := 0; i < n; i++ {
		server[i] = make([]int, m)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Scan(&server[i][j])
		}
	}
	maxCount := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			maxCount = max(maxCount, dfs可以组成网络的服务器(i, j, n, m, 0, server))
		}
	}
	fmt.Println(maxCount)
}
func dfs可以组成网络的服务器(i, j, n, m, count int, server [][]int) int {
	if i < 0 || i >= n || j < 0 || j >= m || server[i][j] == 0 {
		return count
	}
	count++
	server[i][j] = 0
	count = dfs可以组成网络的服务器(i-1, j, n, m, count, server)
	count = dfs可以组成网络的服务器(i+1, j, n, m, count, server)
	count = dfs可以组成网络的服务器(i, j-1, n, m, count, server)
	count = dfs可以组成网络的服务器(i, j+1, n, m, count, server)
	return count
}

// 有一棵二叉树，每个节点由一个大写字母标识(最多26个节点）。现有两组字母，分别表示后序遍历（左孩子->右孩子->父节点）和中序遍历（左孩子->父节点->右孩子）的结果，请输出层次遍历的结果。
// 输入描述：
// 二叉树后续和中序遍历的结果
// 输出描述：
// 二叉树层序遍历的结果
// 总结：1、广度优先遍历需要用队列来实现。2、深度优先需要用递归
type treeNode struct {
	name  string
	left  *treeNode
	right *treeNode
}

func 二叉树的广度优先遍历() {
	var lastorder, inorder string
	fmt.Scan(&lastorder, &inorder)

	root := treeBuildInAndLast(lastorder, inorder)
	bfs二叉树的广度优先遍历(root)
}
func treeBuildInAndLast(lastorder string, inorder string) *treeNode {
	if len(lastorder) == 0 {
		return nil
	}
	//后续遍历最后一个值就是根节点
	root := &treeNode{
		name:  string(lastorder[len(lastorder)-1]),
		left:  nil,
		right: nil,
	}
	rootIndex := -1
	for i, v := range inorder {
		if uint8(v) == lastorder[len(lastorder)-1] {
			rootIndex = i
		}
	}
	root.left = treeBuildInAndLast(lastorder[:rootIndex], inorder[:rootIndex])
	root.right = treeBuildInAndLast(lastorder[rootIndex:len(lastorder)-1], inorder[rootIndex+1:])
	return root
}
func bfs二叉树的广度优先遍历(root *treeNode) {
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		node := queue.Remove(queue.Front()).(*treeNode)
		fmt.Printf("%s", node.name)
		if node.left != nil {
			queue.PushBack(node.left)
		}
		if node.right != nil {
			queue.PushBack(node.right)
		}
	}
}

// 一个局域网内有很多台电脑，分别标注为 0 ~ N-1 的数字。相连接的电脑距离不一样，所以感染时间不一样，感染时间用 t 表示。
// 其中网络内一台电脑被病毒感染，求其感染网络内所有的电脑最少需要多长时间。如果最后有电脑不会感染，则返回-1。
// 给定一个数组 times 表示一台电脑把相邻电脑感染所用的时间。
// 如图：path[i] = {i, j, t} 表示：电脑 i->j，电脑 i 上的病毒感染 j，需要时间 t。
// 输入描述
// 第一行输入一个整数N，表示局域网内电脑个数 N，1<= N<= 200 ；
// 第二行输入一个整数M, 表示有 M 条网络连接；
// 接下来M行, 每行输入为 i,j,t 。表示电脑 i 感染电脑 j 需要时间t。(1 <= i, j <= N)
// 最后一行为病毒所在的电脑编号。
// 4
// 3
// 2 1 1
// 2 3 1
// 3 4 1
// 2
// 输出描述
// 输出最少需要多少时间才能感染全部电脑，如果不存在输出 -1
// 2
// 广度搜索
type network struct {
	index int //病毒电脑编号
	num   int //感染批次
	time  int //感染所需时间
}

func 电脑病毒感染() {
	//电脑台数
	var N int
	fmt.Scan(&N)
	//网络连接数
	var M int
	fmt.Scan(&M)
	var networks = make([][]int, M)
	for i := 0; i < M; i++ {
		networks[i] = make([]int, 3)
	}
	//具体连接
	for i := 0; i < M; i++ {
		for j := 0; j < 3; j++ {
			fmt.Scan(&networks[i][j])
		}
	}
	//病毒所在电脑编号
	var markIndex int
	fmt.Scan(&markIndex)
	//标记已感染病毒的电脑编号
	markSlice := map[int]int{}
	//感染总花费时间
	times := 0
	preNum := 0
	minTime := 0
	queue := list.New()
	queue.PushBack(network{index: markIndex, num: preNum, time: minTime})
	for queue.Len() > 0 {
		node := queue.Remove(queue.Front()).(network)
		markSlice[node.index] = node.index
		if preNum == node.num {
			minTime = min(minTime, node.time)
		} else {
			times += minTime
			minTime = node.time
			preNum = node.num
		}
		num := preNum + 1
		for i := 0; i < M; i++ {
			if networks[i][0] == node.index {
				queue.PushBack(network{index: networks[i][1], num: num, time: networks[i][2]})
			}
		}
	}
	times += minTime
	if len(markSlice) < N {
		fmt.Println(-1)
		return
	}
	fmt.Println(times)
}

// 现有若干个会议，所有会议共享一个会议室，用数组表示各个会议的开始时间和结束时间，
// 格式为: [[会议1开始时间,会议1结束时间],[会议2开始时间,会议2结束时间]] 请计算会议室占用时间段。
// 输入描述
// [[会议1开始时间,会议1结束时间],[会议2开始时间,会议2结束时间]]
// 备注:
// 会议个数范围: [1,100]
// 会议室时间段: [1,24]
// 输出描述
// 输出格式预输入一致,具体请看用例。
// [[会议开始时间，会议结束时间]，[会议开始时间，会议结束时间]
func 会议室占用时间段() {
	reader := bufio.NewReader(os.Stdin)
	str, _ := reader.ReadString('\n')
	str = strings.TrimSuffix(str, "]\n")
	str = strings.TrimPrefix(str, "[")
	strSlice := strings.Split(str, "],[")
	dealSlice := [][2]int64{}
	for _, v := range strSlice {
		v = strings.TrimPrefix(v, "[")
		v = strings.TrimSuffix(v, "]")
		vSlice := strings.Split(v, ",")
		start, _ := strconv.ParseInt(vSlice[0], 10, 64)
		end, _ := strconv.ParseInt(vSlice[1], 10, 64)
		dealSlice = append(dealSlice, [2]int64{start, end})
	}
	resSlice := [][2]int64{}
	markIndex := []int{}
	for i := 0; i < len(dealSlice); i++ {
		if len(markIndex) > 0 && slices.Contains(markIndex, i) {
			continue
		}
		leftMark := dealSlice[i][0]
		rightMark := dealSlice[i][1]
		for j := i + 1; j < len(dealSlice); j++ {
			if len(markIndex) > 0 && slices.Contains(markIndex, i) {
				continue
			}
			left := dealSlice[j][0]
			right := dealSlice[j][1]
			if leftMark <= left && rightMark >= right {
				//包含
			} else if leftMark >= left && rightMark <= right {
				//被包含
				leftMark = left
				rightMark = right
				markIndex = append(markIndex, j)
			} else if leftMark >= left && rightMark >= right && leftMark <= right {
				//交叉
				leftMark = left
				markIndex = append(markIndex, j)
			} else if leftMark <= left && rightMark <= right && rightMark >= left {
				//交叉
				rightMark = right
				markIndex = append(markIndex, j)
			}
		}
		resSlice = append(resSlice, [2]int64{leftMark, rightMark})
	}

	//fmt.Println(resSlice)
	res := "["
	for _, v := range resSlice {
		res += fmt.Sprintf("[%d,%d],", v[0], v[1])
	}
	res = strings.TrimSuffix(res, ",") + "]"
	fmt.Println(res)
}

// 部门在进行需求开发时需要进行人力安排。当前部门需要完成N个需求，需求用requirements[]表示，requirements[i]表示第i个需求的工作量大小，单位：人月。
// 这部分需求需要在M个月内完成开发，进行人力安排后每个月的人力是固定的。
// 目前要求每个月最多有2个需求开发，并且每个月需要完成的需求不能超过部门人力。请帮部门评估在满足需求开发进度的情况下，每个月需要的最小人力是多少？
// 输入描述：
// 输入为M和requirements，M表示需求开发时间要求，requirements表示每个需求工作量大小，N为requirements长度
// 输出描述：
// 对于每一组测试数据，输出部门需要人力需求，行末无多余的空格
func 部门人力分配() {

}

// 马是象棋(包括中国象棋只和国际象棋）中的棋子，走法是每步直一格再斜一格，即先横着或直着走一格，然后再斜着走一个对角线，可进可退，可越过河界，俗称马走 “日“ 字。
// 给项m行n列的棋盘(网格图)，棋盘上只有象棋中的棋子“马”，并目每个棋子有等级之分，等级为K的马可以跳1~k步(走的方式与象棋中“马”的规则一样，不可以超出棋盘位置)，问是否能将所有马跳到同一位置，如果存在，输出最少需要的总步数(每匹马的步数相加) ，不存在则输出-1。
// **注:**允许不同的马在跳的过程中跳到同一位置，坐标为(x,y)的马跳一次可以跳到到坐标为(x+1,y+2),(x+1,y-2),(x+2,y+1),(x+2,y-1). (x-1,y+2),(x-1,y-2),(x-2,y+1),(x-2,y-1),的格点上，但是不可以超出棋盘范围。
// 输入描述
// 第一行输入m,n代表m行n列的网格图棋盘(1 <= m,n <= 25);
// 接下来输入m行n列的网格图棋盘，如果第i行,第j列的元素为 “.” 代表此格点没有棋子，如果为数字k (1<= k <=9)，代表此格点存在等级为的“马”。
// 输出描述
// 输出最少需要的总步数 (每匹马的步数相加)，不存在则输出-1。
type place struct {
	horseNum int //能到当前位置马的数量
	stepNum  int //到当前位置的总步数
}
type horse struct {
	x, y, k int
}

func 跳马() {
	var m, n int
	fmt.Scan(&m, &n)
	chessboard := make([][]string, m)
	count := make([][]place, m)
	for i := 0; i < m; i++ {
		chessboard[i] = make([]string, n)
		count[i] = make([]place, n)
	}
	reader := bufio.NewReader(os.Stdin)
	for i := 0; i < m; i++ {
		str, _ := reader.ReadString('\n')
		str = strings.TrimSuffix(str, "\n")
		for k, v := range str {
			chessboard[i][k] = string(v)
		}
	}

	N := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if chessboard[i][j] != "." {
				k, err := strconv.ParseInt(chessboard[i][j], 10, 64)
				if err != nil {
					continue
				}
				N++
				bfs跳马(int(k), i, j, m, n, count)
			}
		}
	}
	minStepNum := math.MaxInt32
	isExist := false
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if count[i][j].horseNum == N {
				minStepNum = min(minStepNum, count[i][j].stepNum)
				isExist = true
			}
		}
	}
	if !isExist {
		fmt.Println(-1)
		return
	}
	fmt.Println(minStepNum)
}

func bfs跳马(k, startX, startY, m, n int, count [][]place) {
	directions := [][2]int{{1, 2}, {1, -2}, {2, 1}, {2, -1}, {-1, 2}, {-1, -2}, {-2, 1}, {-2, -1}}
	visit := make([][]bool, m)
	for i := 0; i < m; i++ {
		visit[i] = make([]bool, n)
	}
	visit[startX][startY] = true
	queue := list.New()
	queue.PushBack(horse{x: startX, y: startY, k: k})
	for queue.Len() > 0 {
		h := queue.Remove(queue.Front()).(horse)
		count[h.x][h.y].horseNum += 1
		count[h.x][h.y].stepNum += (k - h.k)
		for _, direction := range directions {
			if h.k == 0 {
				break
			}
			if h.x+direction[0] >= 0 && h.x+direction[0] < m && h.y+direction[1] >= 0 && h.y+direction[1] < n && visit[h.x+direction[0]][h.y+direction[1]] == false {
				visit[h.x+direction[0]][h.y+direction[1]] = true
				queue.PushBack(horse{x: h.x + direction[0], y: h.y + direction[1], k: h.k - 1})
			}
		}
	}
}

// 在某个项目中有多个任务（用 tasks 数组表示）需要您进行处理，其中 tasks[i] = [si, ei]，你可以在 si <= day <= ei 中的任意一天处理该任务。请返回你可以处理的最大任务数。
// 注：一天可以完成一个任务的处理。
// 输入描述：
// 第一行为任务数量 n，1 <= n <= 100000。后面 n 行表示各个任务的开始时间和终止时间，用 si 和 ei 表示，1 <= si <= ei <= 100000。
// 输出描述：
// 输出为一个整数，表示可以处理的最大任务数。
func 任务处理() {
	var n int
	fmt.Scan(&n)
	tasks := make([][]int, n)
	for i := 0; i < n; i++ {
		tasks[i] = make([]int, 2)
	}
	maxTime := 0
	for i := 0; i < n; i++ {
		fmt.Scan(&tasks[i][0], &tasks[i][1])
		maxTime = max(maxTime, tasks[i][1])
	}
	count := 0
	indexs := []int{}
	for i := 1; i <= maxTime; i++ {
		if len(indexs) == n {
			break
		}
		minRightValue := 100000
		index := -1
		for j := 0; j < n; j++ {
			if slices.Contains(indexs, j) {
				continue
			}
			if tasks[j][0] <= i && i <= tasks[j][1] {
				if minRightValue > tasks[j][0] {
					minRightValue = tasks[j][0]
					index = j
				}
			}
		}
		if index != -1 {
			indexs = append(indexs, index)
			count++
		}
	}
	fmt.Println(count)
}

// 有一个考古学家发现一个石碑，但是很可惜，发现时其已经断成多段，原地发现n个断口整齐的石碑碎片。为了破解石碑内容，考古学家希望有程序能帮忙计算复原后的石碑文字组合数，你能帮忙吗？
// 输入
// 第一行输入n,n表示石碑碎片的个数
// 第二行依次输入石碑碎片上的文字内容，共有n组
// 输出
// 输出石碑文字的组合（升序），行末无多余空格
func 考古学家() {
	var n int
	fmt.Scan(&n)
	strSlice := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&strSlice[i])
	}
	results := []string{}
	collectStr(strSlice, "", &results)
	sort.Strings(results)
	for _, v := range results {
		fmt.Println(v)
	}
}
func collectStr(strSlice []string, result string, results *[]string) {
	if len(strSlice) == 0 {
		*results = append(*results, result)
		return
	}
	for k, v := range strSlice {
		slice := []string{}
		if k == 0 {
			slice = append(slice, strSlice[k+1:]...)
		} else {
			slice = append(slice, strSlice[:k]...)
			slice = append(slice, strSlice[k+1:]...)
		}
		collectStr(slice, result+v, results)
	}
}

// Wonderland是小王居住地一家很受欢迎的游乐园。Wonderland目前有4种售票方式，分别为一日票（1天）、三日票（3天）、周票（7天）和月票（30天）。
// 每种售票方式的价格由一个数组给出，每种票据在票面时限内可以无限制地进行游玩。例如：
// 小王在第10日买了一张三日票，小王可以在第10日、第11日和第12日进行无限制地游玩。
// 小王计划在接下来一年多次游玩该游乐园。小王计划地游玩日期将由一个数组给出。
// 现在，请您根据给出地售票价格数组和小王计划游玩日期数组，返回游玩计划所需要地最低消费。
// 输入描述:
// 输入为2个数组 售票价格数组为costs，costs.length=4,默认顺序为一日票、三日票、周票和月票。 小王计划游玩日期数组为days，1<=days.length<=365,1<=days[i]<=365,默认顺序为升序。
// 输出描述:
// 完成游玩计划的最低消费 备注: 样例说明: 根据售票价格数组和游玩日期数组给出的信息，发现每次去玩的时候买一张一日票是最省钱的，所以小王会买8张一日票，每张5元，最低花费是40元。
func wonderland1() {
	costs := make([]int, 4)
	for i := 0; i < 4; i++ {
		fmt.Scan(&costs[i])
	}
	reader := bufio.NewReader(os.Stdin)
	daysStr, _ := reader.ReadString('\n')
	daysStr = strings.TrimSuffix(daysStr, "\n")
	days := strings.Split(daysStr, " ")
	daysAll := []int{}
	for _, day := range days {
		d, _ := strconv.Atoi(day)
		daysAll = append(daysAll, d)
	}
	day3to1 := costs[1] / costs[0]
	day7to1 := costs[2] / costs[0]
	day30to1 := costs[3] / costs[0]
	days30 := getMaxDays(&daysAll, 30, day30to1)
	days7 := getMaxDays(&daysAll, 7, day7to1)
	days3 := getMaxDays(&daysAll, 3, day3to1)

	fmt.Println(len(days30)*costs[3] + len(days7)*costs[2] + len(days3)*costs[1] + len(daysAll)*costs[0])
}

// daysAll：计划游玩日期数组
// dayNum：售票类型
// minDays：售票类型对应的回本所需最少天数
func getMaxDays(daysAll *[]int, dayNum int, minDays int) [][]int {
	index := 0
	count := 1
	maxCount := 1
	days := [][]int{}
	for j := index + 1; j < len(*daysAll); j++ {
		count++
		if (*daysAll)[j]-(*daysAll)[index] < dayNum {
			if count > minDays {
				if maxCount < count {
					maxCount = count
				}
			}
		} else {
			if maxCount > 1 {
				days = append(days, (*daysAll)[index:maxCount])
			}
			maxCount = 1
			count = 1
			index = j
		}
	}
	if len(days) > 0 {
		res := []int{}
		deleteIndex := []int{}
		for k, v := range *daysAll {
			for i := 0; i < len(days); i++ {
				for j := 0; j < len(days[i]); j++ {
					if days[i][j] == v {
						deleteIndex = append(deleteIndex, k)
					}
				}
			}
		}

		for k, v := range *daysAll {
			if !slices.Contains(deleteIndex, k) {
				res = append(res, v)
			}
		}

		*daysAll = res
	}
	return days
}

// 这道题用动态规划的方法简单很多
func wonderland2() {
	costs := make([]int, 4)
	for i := 0; i < 4; i++ {
		fmt.Scan(&costs[i])
	}
	reader := bufio.NewReader(os.Stdin)
	daysStr, _ := reader.ReadString('\n')
	daysStr = strings.TrimSuffix(daysStr, "\n")
	days := strings.Split(daysStr, " ")
	daysAll := []int{}
	for _, day := range days {
		d, _ := strconv.Atoi(day)
		daysAll = append(daysAll, d)
	}

	maxDay := daysAll[len(daysAll)-1]
	dp := make([]int, 366)
	for i := 1; i <= maxDay; i++ {
		if slices.Contains(daysAll, i) {
			minMoney := math.MaxInt
			if i-1 >= 0 {
				minMoney = min(minMoney, dp[i-1]+costs[0])
			}
			if i-3 >= 0 {
				minMoney = min(minMoney, dp[i-3]+costs[1])
			}
			if i-7 >= 0 {
				minMoney = min(minMoney, dp[i-7]+costs[2])
			}
			if i-30 >= 0 {
				minMoney = min(minMoney, dp[i-30]+costs[3])
			}
			dp[i] = minMoney
		} else {
			dp[i] = dp[i-1]
		}
	}

	fmt.Println(dp[maxDay])
}

// 小华和小为是很要好的朋友，他们约定周末一起吃饭。
// 通过手机交流，他们在地图上选择了多个聚餐地点（由于自然地形等原因，部分聚餐地点不可达），求小华和小为都能到达的聚餐地点有多少个？
// 输入描述
// 第一行输入m和n，m代表地图的长度，n代表地图的宽度。
// 第二行开始具体输入地图信息，地图信息包含：
// 0 为通畅的道路
// 1 为障碍物（且仅1为障碍物）
// 2 为小华或者小为，地图中必定有且仅有2个 （非障碍物）
// 3 为被选中的聚餐地点（非障碍物）
// 输出描述
// 可以被两方都到达的聚餐地点数量，行末无空格。
type point struct {
	x, y int
}

func 欢乐的周末() {
	var m, n int
	fmt.Scan(&m, &n)
	worldMap := make([][]int, m)
	for i := 0; i < m; i++ {
		worldMap[i] = make([]int, n)
	}
	huawei := [][2]int{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			fmt.Scan(&worldMap[i][j])
			if worldMap[i][j] == 2 {
				huawei = append(huawei, [2]int{i, j})
			}
		}
	}
	count := math.MaxInt32
	count = min(count, bfs欢乐的周末(huawei[0][0], huawei[0][1], m, n, worldMap))
	count = min(count, bfs欢乐的周末(huawei[1][0], huawei[1][1], m, n, worldMap))
	fmt.Println(count)
}
func bfs欢乐的周末(startX, startY, m, n int, worldMap [][]int) int {
	directions := [][]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}
	visit := make([][]bool, m)
	for i := 0; i < m; i++ {
		visit[i] = make([]bool, n)
	}
	count := 0
	queue := list.New()
	queue.PushBack(point{
		x: startX,
		y: startY,
	})
	visit[startX][startY] = true
	for queue.Len() > 0 {
		p := queue.Remove(queue.Front()).(point)
		if worldMap[p.x][p.y] == 3 {
			count++
		}
		for _, d := range directions {
			if p.x+d[0] >= 0 && p.x+d[0] < m && p.y+d[1] >= 0 && p.y+d[1] < n && !visit[p.x+d[0]][p.y+d[1]] && worldMap[p.x+d[0]][p.y+d[1]] != 1 {
				queue.PushBack(point{
					x: p.x + d[0],
					y: p.y + d[1],
				})
				visit[p.x+d[0]][p.y+d[1]] = true
			}
		}
	}
	return count
}

// 一个应用启动时，会有多个初始化任务需要执行，并且任务之间有依赖关系，例如A任务依赖B任务，那么必须在B任务执行完成之后，才能开始执行A任务。
// 现在给出多条任务依赖关系的规则，请输入任务的顺序执行序列，规则采用贪婪策略，即一个任务如果没有依赖的任务，则立刻开始执行，如果同时有多个任务要执行，则根据任务名称字母顺序排序。
// 输入描述
// 输入参数每个元素都表示任意两个任务之间的依赖关系，输入参数中符号"->"表示依赖方向，例如：
// A->B：表示A依赖B
// 多个依赖之间用单个空格分隔
// 输出描述
// 输出排序后的启动任务列表，多个任务之间用单个空格分隔
func 启动多任务排序() {
	reader := bufio.NewReader(os.Stdin)
	str, _ := reader.ReadString('\n')
	str = strings.TrimSuffix(str, "\n")
	strSlice := strings.Split(str, " ")
	//元素权重表
	strWeightMap := map[string]int{}
	for _, v := range strSlice {
		vString := strings.Split(v, "->")
		_, ok := strWeightMap[vString[1]]
		if !ok {
			strWeightMap[vString[1]] = 0
		}

		if value, ok := strWeightMap[vString[0]]; ok {
			strWeightMap[vString[0]] = max(value, strWeightMap[vString[1]]+1)
		} else {
			strWeightMap[vString[0]] = strWeightMap[vString[1]] + 1
		}
	}

	maxWeight := 0
	for _, v := range strWeightMap {
		if v > maxWeight {
			maxWeight = v
		}
	}
	for i := 0; i <= maxWeight; i++ {
		slice := []string{}
		for k, v := range strWeightMap {
			if v == i {
				slice = append(slice, k)
			}
		}
		sort.Strings(slice)
		for j := 0; j < len(slice); j++ {
			fmt.Printf("%s ", slice[j])
		}
	}
}

// 有一种特殊的加密算法，明文为一段数字串，经过密码本查找转换，生成另一段密文数字串。
// 规则如下：
// 明文为一段数字串由 0\~9 组成
// 密码本为数字 0\~9 组成的二维数组
// 需要按明文串的数字顺序在密码本里找到同样的数字串，密码本里的数字串是由相邻的单元格数字组成，上下和左右是相邻的，注意：对角线不相邻，同一个单元格的数字不能重复使用。
// 每一位明文对应密文即为密码本中找到的单元格所在的行和列序号（序号从0开始）组成的两个数宇。
// 如明文第 i 位 Data[i] 对应密码本单元格为 Bookx，则明文第 i 位对应的密文为X Y，X和Y之间用空格隔开。
// 如果有多条密文，返回字符序最小的密文。
// 如果密码本无法匹配，返回"error"。
// 请你设计这个加密程序。
// 输入描述
// 第一行输入 1 个正整数 N，代表明文的长度（1 ≤ N ≤ 200）
// 第二行输入 N 个明文组成的序列 Data[i]（0 ≤ Data[i] ≤ 9）
// 第三行输入 1 个正整数 M，代表密文的长度
// 接下来 M 行，每行 M 个数，代表密文矩阵
// 输出描述
// 输出字典序最小密文，如果无法匹配，输出"error"
func 特殊的加密算法() {
	var N int
	fmt.Scan(&N)
	data := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&data[i])
	}
	var M int
	fmt.Scan(&M)
	book := make([][]int, M)
	visited := make([][]bool, M)
	for i := 0; i < M; i++ {
		book[i] = make([]int, M)
		visited[i] = make([]bool, M)
	}
	for i := 0; i < M; i++ {
		for j := 0; j < M; j++ {
			fmt.Scan(&book[i][j])
		}
	}

	minPath := ""
	for i := 0; i < M; i++ {
		for j := 0; j < M; j++ {
			if data[0] == book[i][j] {
				path := ""
				paths := []string{}
				dfs特殊的加密算法(i, j, M, 0, data, book, visited, path, &paths)
				if len(paths) > 0 {
					sort.Strings(paths)
					if minPath == "" {
						minPath = paths[0]
					} else {
						minPath = min(minPath, paths[0])
					}
				}

			}
		}
	}
	if len(minPath) == 0 {
		fmt.Println("error")
		return
	}
	fmt.Println(minPath)
}
func dfs特殊的加密算法(startX, startY, m, index int, data []int, book [][]int, visited [][]bool, path string, paths *[]string) {
	if startX >= m || startX < 0 || startY >= m || startY < 0 || visited[startX][startY] || index >= len(data) || data[index] != book[startX][startY] {
		return
	}
	path = path + fmt.Sprintf("%d %d ", startX, startY)
	visited[startX][startY] = true
	index++
	if index == len(data) {
		*paths = append(*paths, path)
	}

	dfs特殊的加密算法(startX-1, startY, m, index, data, book, visited, path, paths)
	dfs特殊的加密算法(startX+1, startY, m, index, data, book, visited, path, paths)
	dfs特殊的加密算法(startX, startY-1, m, index, data, book, visited, path, paths)
	dfs特殊的加密算法(startX, startY+1, m, index, data, book, visited, path, paths)
	return
}

// 有一个文件, 包含一定规则写作的文本, 请统计文件中包含的文本数量
// 规则如下
// 1.文本以";"分隔，最后一条可以没有";"，但空文本不能算语句，比如"COMMAND A; ;"只能算一条语句.
// 注意, 无字符/空白字符/制表符都算作"空"文本
// 2. 文本可以跨行, 比如下面, 是一条文本, 而不是三条
// COMMAND A
// AND
// COMMAND B;
// 3. 文本支持字符串, 字符串为成对的单引号(')或者成对的双引号("), 字符串可能出现用转义字符(\)处理的单双引号(比如"your input is: \"")和转义字符本身, 比如
// COMMAND A "Say \"hello\"";
// 4. 支持注释, 可以出现在字符串之外的任意位置, 注释以"--"开头, 到换行结束, 比如
// COMMAND A; -- this is comment
// COMMAND -- comment
// A AND COMMAND B;
// 注意, 字符串内的"--", 不是注释
// 输入描述:
// 文本文件
// 输出描述:
// 包含的文本数量
func 统计文本数量() {

}

// 快递公司每日早晨，给每位快递员推送需要送到客户手中的快递以及路线信息，快递员自己又查找了一些客户与客户之间的路线距离信息，请你依据这些信息，给快递员设计一条最例短路径， 告诉他最短路径的距离。
// 注意：
// 不限制快递包裹送到客户手中的顺序，但必须保证都送到客户手中
// 用例保证一定存在投递站到每位客户之间的路线，但不保证客户与客户之间有路线，客户位置及投递站均允许多次经过。
// 所有快递送完后，快递员需回到投递站
// 输入描述
// 首行输入两个正整数n,m。
// 接下面n行，输入快递公司发布的客户快递信息，
// 格式为：客户id 投递站到客户之间的距离distance
// 再接下来的m行，是快递员自行查找的客户与客户之间的距离信息，
// 格式为：客户1id 客户2id distance
// 在每行数据中，数据与数据之间均以单个空格分割
// 输出描述
// 最短路径距离，如无法找到，请输出 -1
func 快递员的烦恼1() {
	var n, m int
	fmt.Scan(&n, &m)
	distances := map[string]int{}
	for i := 0; i < n; i++ {
		pName := ""
		distance := 0
		fmt.Scan(&pName, &distance)
		distances[pName] = distance
	}
	ptopDistance := map[string]map[string]int{}
	for i := 0; i < m; i++ {
		p1, p2, distance := "", "", 0
		fmt.Scan(&p1, &p2, &distance)
		if _, ok := ptopDistance[p1]; ok {
			ptopDistance[p1][p2] = distance
		} else {
			ptopDistance[p1] = map[string]int{p2: distance}
		}
		if _, ok := ptopDistance[p2]; ok {
			ptopDistance[p2][p1] = distance
		} else {
			ptopDistance[p2] = map[string]int{p1: distance}
		}
	}

	//动态规划，从快递站到各个客户的最小距离
	dp := make([]int, n+1)
	dp[1] = distances["1"] * 2
	for i := 2; i <= n; i++ {
		p1 := strconv.FormatInt(int64(i), 10)
		minDis := distances[p1] * 2
		if dis, ok := ptopDistance[p1]; ok {
			for p2, distance := range dis {
				if p2 < p1 {
					minDis = min(minDis, distance+distances[p1]-distances[p2])
				}
				minDis = min(minDis, 2*distance)
			}
		}
		dp[i] = dp[i-1] + minDis
	}
	//fmt.Println(dp)
	fmt.Println(dp[n])
}

func 快递员的烦恼2() {
	var n, m int
	fmt.Scan(&n, &m)
	distance := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		distance[i] = make([]int, n+1)
	}
	for i := 0; i <= n; i++ {
		for j := 0; j <= n; j++ {
			if i != j {
				distance[i][j] = math.MaxInt32
			}
		}
	}

	for i := 0; i < n; i++ {
		pName, dist := 0, 0
		fmt.Scan(&pName, &dist)
		distance[pName][0] = dist
		distance[0][pName] = dist
	}
	for i := 0; i < m; i++ {
		j, k, dis := 0, 0, 0
		fmt.Scan(&j, &k, &dis)
		distance[j][k] = dis
		distance[k][j] = dis
	}

	//Floyd-Warshall算法预处理任意两点间的最短路径
	for i := 0; i <= n; i++ {
		for j := 0; j <= n; j++ {
			for z := 0; z <= n; z++ {
				if distance[i][j] > distance[i][z]+distance[z][j] {
					distance[i][j] = distance[i][z] + distance[z][j]
				}
			}
		}
	}

	minDis := 0
	for i := 1; i <= n; i++ {
		if distance[i-1][i] == math.MaxInt32 {
			fmt.Println(-1)
			return
		}
		minDis += distance[i-1][i]
	}
	minDis += distance[0][n]
	fmt.Println(minDis)
}

// 开头和结尾都是元音字母（aeiouAEIOU）的字符串为元音字符串，其中混杂的非元音字母数量为其瑕疵度。比如:
// “a” 、 “aa”是元音字符串，其瑕疵度都为0
// “aiur”不是元音字符串（结尾不是元音字符）
// “abira”是元音字符串，其瑕疵度为2
// 给定一个字符串，请找出指定瑕疵度的最长元音字符子串，并输出其长度，如果找不到满足条件的元音字符子串，输出0。
// 子串：字符串中任意个连续的字符组成的子序列称为该字符串的子串。
// 输入描述
// 首行输入是一个整数，表示预期的瑕疵度flaw，取值范围[0, 65535]。
// 接下来一行是一个仅由字符a-z和A-Z组成的字符串，字符串长度(0, 65535]。
// 输出描述
// 输出为一个整数，代表满足条件的元音字符子串的长度。
func 最长的指定瑕疵度的元音子串() {
	var n int
	fmt.Scan(&n)
	var str string
	fmt.Scan(&str)
	y := []string{"a", "e", "i", "o", "u", "A", "E", "I", "O", "U"}
	maxLength := 0
	for i := 0; i < len(str); i++ {
		m := 0
		if !slices.Contains(y, string(str[i])) {
			continue
		}
		for j := i + 1; j < len(str); j++ {
			if !slices.Contains(y, string(str[j])) {
				m++
				if m > n {
					break
				}
			} else {
				if m == n {
					maxLength = max(maxLength, j-i+1)
				}
			}
		}
	}

	fmt.Println(maxLength)
}

// 请实现一个简易内存池,根据请求命令完成内存分配和释放。
// 内存池支持两种操作命令，REQUEST和RELEASE，其格式为：
// REQUEST=请求的内存大小 表示请求分配指定大小内存，如果分配成功，返回分配到的内存首地址；如果内存不足，或指定的大小为0，则输出error。
// RELEASE=释放的内存首地址 表示释放掉之前分配的内存，释放成功无需输出，如果释放不存在的首地址则输出error。
// 注意：
// 内存池总大小为100字节。
// 内存池地址分配必须是连续内存，并优先从低地址分配。
// 内存释放后可被再次分配，已释放的内存在空闲时不能被二次释放。
// 不会释放已申请的内存块的中间地址。
// 释放操作只是针对首地址所对应的单个内存块进行操作，不会影响其它内存块。
// 输入描述
// 首行为整数 N , 表示操作命令的个数，取值范围：0 < N <= 100。
// 接下来的N行, 每行将给出一个操作命令，操作命令和参数之间用 “=”分割。
// 输出描述
// 请求分配指定大小内存时，如果分配成功，返回分配到的内存首地址；如果内存不足，或指定的大小为0，则输出error
// 释放掉之前分配的内存时，释放成功无需输出，如果释放不存在的首地址则输出error。
func 简易内存池() {
	var N int
	fmt.Scan(&N)
	res := []string{}
	var j int64 = 0
	for i := 0; i < N; i++ {
		var str string
		fmt.Scanln(&str)
		strSlice := strings.Split(str, "=")
		v, _ := strconv.ParseInt(strSlice[1], 10, 64)
		if v == 0 {
			//fmt.Println("error")
			res = append(res, "error")
			continue
		}
		if strSlice[0] == "REQUEST" {
			if j+v > 100 {
				res = append(res, "error")
			} else {
				res = append(res, strconv.FormatInt(j, 10))
				j += v
			}
		} else if strSlice[0] == "RELEASE" {
			if j-v < 0 {
				res = append(res, "error")
			} else {
				res = append(res, strconv.FormatInt(j, 10))
				j -= v
			}
		}
	}

	for _, v := range res {
		fmt.Println(v)
	}
}

// 给定一个表达式，求其分数计算结果。
// 表达式的限制如下：
// 所有的输入数字皆为正整数（包括0）
// 仅支持四则运算（+-*/）和括号
// 结果为整数或分数，分数必须化为最简格式（比如6，3/4，7/8，90/7）
// 除数可能为0，如果遇到这种情况，直接输出"ERROR"
// 输入和最终计算结果中的数字都不会超出整型范围
// 用例输入一定合法，不会出现括号匹配的情况
// 输入描述
// 字符串格式的表达式，仅支持+-*/，数字可能超过两位，可能带有空格，没有负数
// 长度小于200个字符
// 输出描述
// 表达式结果，以最简格式表达
// 如果结果为整数，那么直接输出整数
// 如果结果为负数，那么分子分母不可再约分，可以为假分数，不可表达为带分数
// 结果可能是负数，符号放在前面
// 1*(3*4/(8-(7+0)))
type Fraction struct {
	ch int //分子
	fa int //分母
}

func 符号运算() {
	var str string
	fmt.Scan(&str)
	numStack := list.New()
	opStack := list.New()
	for _, v := range str {
		if slices.Contains([]string{"+", "-", "*", "/", "(", ")"}, string(v)) {
			if slices.Contains([]string{"(", ")"}, string(v)) {
				continue
			}
			opStack.PushBack(string(v))
		} else {
			value, _ := strconv.ParseInt(string(v), 10, 64)
			numStack.PushBack(int(value))
		}
	}
	newFraction := Fraction{}
	n := 1
	for opStack.Len() > 0 {
		op := opStack.Remove(opStack.Back()).(string)
		v1 := numStack.Remove(numStack.Back()).(int)
		v2 := 0
		if n == 1 {
			v2 = numStack.Remove(numStack.Back()).(int)
		}
		switch op {
		case "+":
			if n == 1 {
				newFraction.ch = v2 + v1
				newFraction.fa = 1
			} else {
				newFraction.ch = newFraction.ch + newFraction.fa*v1
			}
		case "-":
			if n == 1 {
				newFraction.ch = v2 - v1
				newFraction.fa = 1
			} else {
				newFraction.ch = newFraction.fa*v1 - newFraction.ch
			}
		case "*":
			if n == 1 {
				newFraction.ch = v2 * v1
				newFraction.fa = 1
			} else {
				newFraction.ch = newFraction.ch * v1
			}
		case "/":
			if n == 1 {
				newFraction.ch = v2
				newFraction.fa = v1
			} else {
				newFraction.fa = newFraction.ch
				newFraction.ch = v1 * newFraction.fa
			}
		}
		n++
	}

	divisor := gcd(newFraction.ch, newFraction.fa)
	if divisor == newFraction.fa {
		fmt.Println(newFraction.ch / divisor)
	} else {
		fmt.Println(fmt.Sprintf("%d/%d", newFraction.ch/divisor, newFraction.fa/divisor))
	}
}

// gcd 计算两个数的最大公约数
func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

// 给定 M（0 < M ≤ 30）个字符（a-z），从x中取出任意字符（每个字符只能用一次）拼接成长度为 N（0 < N ≤ 5）的字符串，
// 要求相同的字符不能相邻，计算出给定的字符列表能拼接出多少种满足条件的字符串，
// 输入非法或者无法拼接出满足条件的字符串则返回0。
// 输入描述
// 给定的字符列表和结果字符串长度，中间使用空格(" ")拼接
// 输出描述
// 满足条件的字符串个数
func 字符串拼接() {
	var str string
	var N int
	fmt.Scan(&str, &N)
	resSlice := []string{}
	for i := 0; i < len(str); i++ {
		dfs字符串拼接(str, N-1, string(str[i]), &resSlice, []int{i})
	}

	uniqueRes := map[string]string{}
	result := []string{}
	for _, v := range resSlice {
		if _, ok := uniqueRes[v]; ok {
			continue
		}
		result = append(result, v)
		uniqueRes[v] = v
	}

	//fmt.Println(result)
	fmt.Println(len(result))
}
func dfs字符串拼接(str string, N int, res string, resSlice *[]string, index []int) {
	if str == "" || N == 0 || len(str) < N {
		if N == 0 {
			*resSlice = append(*resSlice, res)
		}
		return
	}
	for i := 0; i < len(str); i++ {
		m := N
		markIndex := index
		result := res
		if slices.Contains(markIndex, i) || str[i] == result[len(result)-1] {
			continue
		}
		result += string(str[i])
		markIndex = append(markIndex, i)
		m--
		dfs字符串拼接(str, m, result, resSlice, markIndex)
	}
}

// 给定一个有向图，图中可能包含有环，图使用二维矩阵表示，每一行的第一列表示起始节点，第二列表示终止节点，如 [0, 1] 表示从 0 到 1 的路径。
// 每个节点用正整数表示。求这个数据的首节点与尾节点，题目给的用例会是一个首节点，但可能存在多个尾节点。同时图中可能含有环。如果图中含有环，返回 [-1]。
// 说明：入度为0是首节点，出度为0是尾节点。
// 输入描述
// 第一行为后续输入的键值对数量N（N ≥ 0）
// 第二行为2N个数字。每两个为一个起点，一个终点.
// 输出描述
// 输出一行头节点和尾节点。如果有多个尾节点，按从小到大的顺序输出
// 备注
// 如果图有环，输出为 -1
// 所有输入均合法，不会出现不配对的数据
type pictureNode struct {
	value int
	deep  int
}

func 查找一个有向网络的头节点和尾节点() {
	var N int
	fmt.Scan(&N)
	m := make([][]int, N)
	for i := 0; i < N; i++ {
		m[i] = make([]int, 2)
		fmt.Scan(&m[i][0], &m[i][1])
	}
	//检查是否有环
	for i := 0; i < N; i++ {
		for j := i + 1; j < N; j++ {
			if m[j][0] == m[i][1] && m[j][1] == m[i][0] {
				fmt.Println(-1)
				return
			}
		}
	}

	resSlice := []pictureNode{}
	maxDeep := 0
	for i := 0; i < N; i++ {
		deep := 0
		one := false
		two := false
		for _, v := range resSlice {
			if v.value == m[i][0] {
				deep = max(deep, v.deep)
				one = true
			}
			if v.value == m[i][1] {
				deep = max(deep, v.deep)
				two = true
			}
		}
		if !one || !two {
			if !one && !two {
				resSlice = append(resSlice, pictureNode{m[i][0], deep})
				deep += 1
				maxDeep = max(maxDeep, deep)
				resSlice = append(resSlice, pictureNode{m[i][1], deep})
			} else {
				if !one {
					deep -= 1
					resSlice = append(resSlice, pictureNode{m[i][0], deep})
				} else {
					deep += 1
					maxDeep = max(maxDeep, deep)
					resSlice = append(resSlice, pictureNode{m[i][1], deep})
				}
			}
		} else {
			for k, v := range resSlice {
				if v.value == m[i][1] {
					deep += 1
					maxDeep = max(maxDeep, deep)
					resSlice[k].deep = deep
					break
				}
			}
		}
	}

	endSlice := []int{}
	startValue := 0
	for _, v := range resSlice {
		if v.deep == 0 {
			startValue = v.value
		}
		if v.deep == maxDeep {
			endSlice = append(endSlice, v.value)
		}
	}
	slices.Sort(endSlice)
	endSlice = append([]int{startValue}, endSlice...)
	for _, v := range endSlice {
		fmt.Printf("%d ", v)
	}
}

// 请设计一个文件缓存系统，该文件缓存系统可以指定缓存的最大值（单位为字节）。
// 文件缓存系统有两种操作：
// 存储文件（put）
// 读取文件（get）
// 操作命令为：
// put fileName fileSize
// get fileName
// 存储文件是把文件放入文件缓存系统中；
// 读取文件是从文件缓存系统中访问已存在，如果文件不存在，则不作任何操作。
// 当缓存空间不足以存放新的文件时，根据规则删除文件，直到剩余空间满足新的文件大小位置，再存放新文件。
// 具体的删除规则为：
// 文件访问过后，会更新文件的最近访问时间和总的访问次数，当缓存不够时，按照第一优先顺序为访问次数从少到多，第二顺序为时间从老到新的方式来删除文件。
// 输入描述
// 第一行为缓存最大值 m（整数，取值范围为 0 < m ≤ 52428800）
// 第二行为文件操作序列个数 n（0 ≤ n ≤ 300000）
// 从第三行起为文件操作序列，每个序列单独一行，文件操作定义为：
// op file\_name file\_size
// file\_name 是文件名，file\_size 是文件大小
// 输出描述
// 输出当前文件缓存中的文件名列表，文件名用英文逗号分隔，按字典顺序排序，如：
// a,c
// 如果文件缓存中没有文件，则输出NONE
// 备注
// 如果新文件的文件名和文件缓存中已有的文件名相同，则不会放在缓存中
// 新的文件第一次存入到文件缓存中时，文件的总访问次数不会变化，文件的最近访问时间会更新到最新时间
// 每次文件访问后，总访问次数加1，最近访问时间更新到最新时间
// 任何两个文件的最近访问时间不会重复
// 文件名不会为空，均为小写字母，最大长度为10
// 缓存空间不足时，不能存放新文件
// 每个文件大小都是大于 0 的整数输入
type fileNode struct {
	fileName   string
	cache      int
	times      int
	visitOrder int
}

func 文件缓存系统() {
	var m int
	var n int
	fmt.Scanln(&m)
	fmt.Scanln(&n)
	fileList := map[string]fileNode{}
	for i := 0; i < n; i++ {
		var a, b string
		var c int
		fmt.Scan(&a)
		if a == "put" {
			fmt.Scan(&b, &c)
			if _, ok := fileList[b]; ok {
				continue
			}
			if m >= c {
				fileList[b] = fileNode{
					fileName:   b,
					cache:      c,
					times:      0,
					visitOrder: 0,
				}
				m -= c
			} else {
				for {
					minTimes := n
					minOrder := n
					fileName := ""
					for _, file := range fileList {
						minTimes = min(minTimes, file.times)
					}
					for name, file := range fileList {
						if file.times == minTimes {
							if file.visitOrder < minOrder {
								fileName = name
							}
						}
					}
					m += fileList[fileName].cache
					delete(fileList, fileName)
					if m >= c {
						fileList[b] = fileNode{
							fileName:   b,
							cache:      c,
							times:      0,
							visitOrder: 0,
						}
						m -= c
						break
					}
				}
			}
		} else if a == "get" {
			fmt.Scan(&b)
			if file, ok := fileList[b]; !ok {
				continue
			} else {
				file.times++
				file.visitOrder = i
				fileList[b] = file
			}
		}
	}
	fileName := []string{}
	for name, _ := range fileList {
		fileName = append(fileName, name)
	}
	if len(fileName) == 0 {
		fmt.Println("NONE")
		return
	}
	sort.Strings(fileName)
	res := strings.Join(fileName, ",")
	fmt.Println(res)
}

type networkNode struct {
	x, y   int
	signal []int
}

func 寻找最优的路测线路() {
	var R, C int
	fmt.Scanln(&R)
	fmt.Scanln(&C)
	network := make([][]int, R)
	for i := 0; i < R; i++ {
		network[i] = make([]int, C)
		for j := 0; j < C; j++ {
			fmt.Scan(&network[i][j])
		}
	}
	fmt.Println(bfs寻找最优的路测线路(0, 0, R-1, C-1, R, C, network))
}

func bfs寻找最优的路测线路(startX, startY, endX, endY, R, C int, network [][]int) int {
	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	visited := make([][]bool, R)
	for i := 0; i < R; i++ {
		visited[i] = make([]bool, C)
	}
	maxValue := -1
	visited[startX][startY] = true
	queue := list.New()
	queue.PushBack(networkNode{0, 0, []int{network[startX][startY]}})
	for queue.Len() > 0 {
		node := queue.Remove(queue.Front()).(networkNode)
		if node.x == endX && node.y == endY {
			sort.Ints(node.signal)
			maxValue = max(maxValue, node.signal[0])
			continue
		}
		for _, direction := range directions {
			x := node.x + direction[0]
			y := node.y + direction[1]
			if x >= 0 && x < R && y >= 0 && y < C && visited[x][y] == false {
				visited[x][y] = true
				queue.PushBack(networkNode{
					x:      x,
					y:      y,
					signal: append(node.signal, network[x][y]),
				})
			}
		}
	}
	return maxValue
}

func 生成哈夫曼树() {
	var N int
	fmt.Scan(&N)
	slice := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&slice[i])
	}
	sort.Ints(slice)

	queue := list.New()
	for _, v := range slice {
		hNode := node{
			value: v,
			left:  nil,
			right: nil,
		}
		queue.PushBack(hNode)
	}
	var root = &node{}
	for queue.Len() > 0 {
		hNode1 := queue.Remove(queue.Front()).(node)
		if root.value != 0 {
			if hNode1.value <= root.value {
				root = &node{
					value: hNode1.value + root.value,
					left:  &hNode1,
					right: root,
				}
			}
		} else {
			if queue.Len() > 0 {
				hNode2 := queue.Remove(queue.Front()).(node)
				root = &node{
					value: hNode1.value + hNode2.value,
					left:  &hNode1,
					right: &hNode2,
				}
			} else {
				root = &hNode1
			}
		}
	}

	inOrder(root)
}
