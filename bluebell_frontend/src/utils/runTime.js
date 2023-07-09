/**
 *  格式化日期
 *  @param {String/Number} date 日期或时间戳
 *  @param {format} 输出形式 如果为空则默认返回的是一个对象,否则返回 YYYY-MM-DD HH:mm:ss，YYYY/MM/DD HH:mm:ss
 * 
 */
function formatDate(date, format) {
	format = format || '';
	let time = new Date(date); // 初始化日期
	let year = time.getFullYear();
	let month = time.getMonth() + 1;
	let day = time.getDate();
	let hour = time.getHours();
	let minute = time.getMinutes();
	let second = time.getSeconds();
	let milliSecond = time.getMilliseconds();
	let resutTimeObject = {
		year,
		month,
		day,
		hour,
		minute,
		second,
		milliSecond
	};
	if (!format) {
		return resutTimeObject;
	} else {
		format = format.trim();
		if (format === 'YYYY-MM-DD hh:mm:ss') {
			return year + '-' + month + '-' + day + ' ' + hour + ':' + minute + ':' + second;
		} else if (format === 'YYYY/MM/DD hh:mm:ss') {
			return year + '/' + month + '/' + day + ' ' + hour + ':' + minute + ':' + second;
		} else if (format === 'YYYY-MM-DD') {
			return year + '-' + month + '-' + day;
		} else if (format === 'YYYY/MM/DD') {
			return year + '/' + month + '/' + day;
		} else if (format === 'hh:mm:ss') {
			return hour + ':' + minute + ':' + second;
		} else if (format === 'hh-mm-ss') {
			return hour + '-' + minute + '-' + second;
		} else {
			return date;
		}
	}
}


/**
 * 定义方法固定位数（位数不够，前面补充自定义数字）
 *  @param {Number} num：被操作数
 *  @param {Number} n： 固定的总位数
 *  @param {Number} num1：位数不足时要补充的数
 *  @returns {String}
 */

function fixedNumber(num, n, num1) {
	num1 = num1 || 0;
	return (Array(n).join(num1) + num).slice(-n);
}


/**
 * 处理时间
 * @param {Number} unitSecond 
 * @param {String} format 
 * @returns 
 */
function handleTime(unitSecond, format) {
	// 创建数组存储 年、日、时、分、秒
	let timeArr = new Array(0, 0, 0, 0, 0);
	// 将秒转换成对应的 年 日 时 分 秒
	let unitYear = 365 * 24 * 60 * 60;
	let unitDay = 24 * 60 * 60;
	let unitHour = 60 * 60;
	let unitMin = 60;
	let unitSec = 0;
	if (!unitSecond) {
		return;
	}
	if ((format ? format.indexOf('Y') > -1 : false) && unitSecond >= unitYear) {
		timeArr[0] = parseInt(unitSecond / unitYear);
		unitSecond %= unitYear;
	}
	if (unitSecond >= unitDay) {
		timeArr[1] = parseInt(unitSecond / unitDay);
		unitSecond %= unitDay;
	}
	if (unitSecond >= unitHour) {
		timeArr[2] = parseInt(unitSecond / unitHour);
		unitSecond %= unitHour;
	}
	if (unitSecond >= unitMin) {
		timeArr[3] = parseInt(unitSecond / unitMin);
		unitSecond %= unitMin;
	}
	if (unitSecond > unitSec) {
		timeArr[4] = unitSecond;
	}
	return timeArr;
}

/**
 * 获取时间
 * @param {Number} year 
 * @param {Number} month 
 * @param {Number} day 
 * @param {Number} hour 
 * @param {Number} minute 
 * @param {Number} second 
 * @returns 
 */
function getTime(year, month, day, hour, minute, second) {
	// 初始化起始时间
	let startTime = Math.round(new Date(Date.UTC(year, month - 1, day, hour, minute, second)).getTime() / 1000);
	// 获取当前时间(中国时区和UTC世界标椎时间相差 8 个小时)
	let nowTime = Math.round((new Date().getTime() + 8 * 60 * 60 * 1000) / 1000);
	return handleTime(nowTime - startTime);
}

/**
 * 格式化运行时间结果
 * @param {Array} runTimeArr 时间数组
 * @param {String} format 格式 
 * @returns 
 */
function handleResult(runTimeArr, format) {
	if (!format) {
		return runTimeArr[1] + '天' + runTimeArr[2] + '时' + runTimeArr[3] + '分' + runTimeArr[4] + '秒';
	} else if (format === 'Y-D h:m:s') {
		return runTimeArr[0] + '年' + runTimeArr[1] + '天' + runTimeArr[2] + '时' + runTimeArr[3] + '分' + runTimeArr[4] + '秒';
	} else if (format === 'D h:m:s') {
		return runTimeArr[1] + '天' + runTimeArr[2] + '时' + runTimeArr[3] + '分' + runTimeArr[4] + '秒';
	}
}

/**
 * 运行时间
 * @param {String/Number/Object} timeStamp 时间戳
 * @param {String/DOMObject} el 展示时间的DOM的选择器
 * @param {String} desc 自定义描述文本
 * @param {Number} year
 * @param {Number} month
 * @param {Number} day
 * @param {Number} hour
 * @param {Number} minute
 * @param {Number} second
 * @param {boolean} flag 返回时间的方式，如果为true，返回带有描述性文本格式的时间模板，如果为false，则返回时间数组，用户可操作返回的时间数组。
 *      
 */
export function runTime({
	el,
	timeStamp,
	desc,
	year,
	month,
	day,
	hour,
	minute,
	second,
	flag = true,
	format
}) {
	desc = desc || '';
	if (timeStamp) {
		let time = formatDate(timeStamp);
		year = time.year;
		month = time.month;
		day = time.day;
		hour = time.hour;
		minute = time.minute;
		second = time.second;
	}
	if (flag && el) {
		let time_wrapper = document.querySelector(el);
		//开始计时
		setInterval(() => {
			let runTimeArr = getTime(year, month, day, hour, minute, second);
			runTimeArr[2] = fixedNumber(runTimeArr[2], 2, 0);
			runTimeArr[3] = fixedNumber(runTimeArr[3], 2, 0);
			runTimeArr[4] = fixedNumber(runTimeArr[4], 2, 0);
			time_wrapper.innerText = desc + handleResult(runTimeArr, format);
		}, 1000);
	} else {
		return getTime(year, month, day, hour, minute, second);
	}
}