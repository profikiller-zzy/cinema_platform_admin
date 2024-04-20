import dayjs from 'dayjs';
export function dateTransition(timestamp, type) {
    // 后端获取到的时间戳为 timestamp
    // 创建一个新的 Date 对象，并将时间戳作为参数传入
    const date = new Date(timestamp);

    // 获取年、月、日、小时、分钟信息
    const year = date.getFullYear();
    const month = (date.getMonth() + 1).toString().padStart(2, "0"); // 月份从 0 开始，需加 1，然后补零
    const day = date.getDate().toString().padStart(2, "0"); // 补零
    const hours = date.getHours().toString().padStart(2, "0"); // 补零
    const minutes = date.getMinutes().toString().padStart(2, "0"); // 补零

    // 将年月日时分拼接成字符串
    if (type === 0) {
        const formattedDateTime = `${year}-${month}-${day} ${hours}:${minutes}`;
        return formattedDateTime
    } else {
        const formattedDateTime = `${year}-${month}-${day}`;
        // 输出格式化后的日期时间字符串
        return formattedDateTime
    }
}