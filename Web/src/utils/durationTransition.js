export function durationTransition(timeDuration) {
     // 将纳秒转换为秒
    const seconds = Math.floor(timeDuration / 1000000000);

    // 计算小时、分钟和剩余秒数
    const hours = Math.floor(seconds / 3600);
    const minutes = Math.floor((seconds % 3600) / 60);
    const remainingSeconds = seconds % 60;

    // 格式化为字符串
    const formattedDurationTime = `${hours}小时${minutes}分钟${remainingSeconds}秒`;

    return formattedDurationTime;
}

export function durationToNanoseconds(durationString) {
  // 使用正则表达式匹配字符串中的小时、分钟和秒
  const match = durationString.match(/(\d+)小时(\d+)分钟(\d+)秒/);
  if (!match) {
    // 如果匹配失败，返回0
    return 0;
  }

  // 将匹配到的小时、分钟和秒转换为数字
  const hours = parseInt(match[1]);
  const minutes = parseInt(match[2]);
  const seconds = parseInt(match[3]);

  // 计算总纳秒数并返回
  return (hours * 3600 + minutes * 60 + seconds) * 1000000000;
}