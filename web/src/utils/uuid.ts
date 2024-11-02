/**
 * UUID生成器
 * @param len 长度 number
 * @param radix 随机数基数 number
 * @returns {string}
 */
export const uuid = (len: number, radix: number = 62) => {
  const chars = '0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz'.split('');
  const uuids = [];
  let i;

  if (len) {
    // Compact form
    for (i = 0; i < len; i++) {
      uuids[i] = chars[Math.floor(Math.random() * radix)];
    }
  } else {
    // rfc4122, version 4 form
    let r;

    // rfc4122 requires these characters
    uuids[8] = uuids[13] = uuids[18] = uuids[23] = '-';
    uuids[14] = '4';

    // Fill in random data.  At i==19 set the high bits of clock sequence as
    // per rfc4122, sec. 4.1.5
    for (i = 0; i < 36; i++) {
      if (!uuids[i]) {
        r = Math.floor(Math.random() * 16);
        uuids[i] = chars[i === 19 ? ((r % 4) % 8) + 8 : r];
      }
    }
  }
  return uuids.join('');
};
