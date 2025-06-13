/**
 * YYYY-MM-DD 形式の日付文字列を「M月D日 (曜)」形式に変換する
 * @param dateString - "YYYY-MM-DD" 形式の日付文字列
 * @returns フォーマットされた日付文字列 (例: "6月13日 (金)")
 */
export const formatToJapaneseDate = (dateString: string): string => {
  try {
    const date = new Date(dateString);
    
    // dateStringが不正な場合に "Invalid Date" となるのを防ぐ
    if (isNaN(date.getTime())) {
      return "無効な日付";
    }

    const month = date.getMonth() + 1;
    const day = date.getDate();
    const dayOfWeek = ["日", "月", "火", "水", "木", "金", "土"][date.getDay()];

    return `${month}月${day}日 (${dayOfWeek})`;
  } catch (error) {
    console.error("Date formatting failed:", error);
    return "日付エラー";
  }
};