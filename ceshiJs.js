function base64decode(str) {
    var base64EncodeChars = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/";
    var base64DecodeChars = new Array(-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 62, -1, -1, -1, 63, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, -1, -1, -1, -1, -1, -1, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, -1, -1, -1, -1, -1, -1, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, -1, -1, -1, -1, -1);
    var c1, c2, c3, c4;
    var i, len, out;
    len = str.length;
    i = 0;
    out = "";
    while (i < len) {
        do {
            c1 = base64DecodeChars[str.charCodeAt(i++) & 255]
        } while (i < len && c1 == -1);
        if (c1 == -1) {
            break
        }
        do {
            c2 = base64DecodeChars[str.charCodeAt(i++) & 255]
        } while (i < len && c2 == -1);
        if (c2 == -1) {
            break
        }
        out += String.fromCharCode((c1 << 2) | ((c2 & 48) >> 4));
        do {
            c3 = str.charCodeAt(i++) & 255;
            if (c3 == 61) {
                return out
            }
            c3 = base64DecodeChars[c3]
        } while (i < len && c3 == -1);
        if (c3 == -1) {
            break
        }
        out += String.fromCharCode(((c2 & 15) << 4) | ((c3 & 60) >> 2));
        do {
            c4 = str.charCodeAt(i++) & 255;
            if (c4 == 61) {
                return out
            }
            c4 = base64DecodeChars[c4]
        } while (i < len && c4 == -1);
        if (c4 == -1) {
            break
        }
        out += String.fromCharCode(((c3 & 3) << 6) | c4)
    }
    return out
};
var ret_classurl = '/mh/2988/';
var comicname = "异皇重生";
var viewid = "945951";
var viewtype = "1";
var viewname = "298混乱的真相";
var photosr = new Array();

packed = "ZXZhbChmdW5jdGlvbihwLGEsYyxrLGUsZCl7ZT1mdW5jdGlvbihjKXtyZXR1cm4gYy50b1N0cmluZygzNil9O2lmKCEnJy5yZXBsYWNlKC9eLyxTdHJpbmcpKXt3aGlsZShjLS0pe2RbZShjKV09a1tjXXx8ZShjKX1rPVtmdW5jdGlvbihlKXtyZXR1cm4gZFtlXX1dO2U9ZnVuY3Rpb24oKXtyZXR1cm4nXFx3Kyd9O2M9MX07d2hpbGUoYy0tKXtpZihrW2NdKXtwPXAucmVwbGFjZShuZXcgUmVnRXhwKCdcXGInK2UoYykrJ1xcYicsJ2cnKSxrW2NdKX19cmV0dXJuIHB9KCdjWzFdPSJlL2IvYS9kL2cvbS5mLzAiO2NbMl09ImUvYi9hL2QvZy9vLmYvMCI7Y1szXT0iZS9iL2EvZC9nL2wuZi8wIjtjWzRdPSJlL2IvYS9kL2cvcC5mLzAiO2NbNV09ImUvYi9hL2QvZy9rLmYvMCI7Y1s2XT0iZS9iL2EvZC9nL2guZi8wIjtjWzddPSJlL2IvYS9kL2cvaS5mLzAiO2NbOF09ImUvYi9hL2QvZy9qLmYvMCI7Y1s5XT0iZS9iL2EvZC9nL24uZi8wIjtjW3FdPSJlL2IvYS9kL2cveS5mLzAiO2NbeF09ImUvYi9hL2QvZy96LmYvMCI7Y1t2XT0iZS9iL2EvZC9nL3cuZi8wIjtjW3JdPSJlL2IvYS9kL2cvcy5mLzAiO2NbdF09ImUvYi9hL2QvZy91LmYvMCI7JywzNiwzNiwnfHx8fHx8fHx8fDA2fDIwMjJ8cGhvdG9zcnwwOHxpbWFnZXMyfGpwZ3wwOXw0NDEwYWI5YjQyfDQ0ZGM3NzRlYTh8NDRlZTNiZjI2Znw0NDA4Y2I4OTY5fDQ0MjNhYzVmMjh8NDQ2OWI5NWY1ZHw0NGRmMDY2MWQyfDQ0NTQwZGQyYTh8NDRhZWEwN2I5MHwxMHwxM3w0NWM0YTU5OTQ3fDE0fDQ1ZTRjYWEzYTF8MTJ8NDU2ZjFiOTAyYnwxMXw0NDlhZDIyNWFmfDQ0YmUxNTFlY2YnLnNwbGl0KCd8JyksMCx7fSkpCg==";
eval(eval(base64decode(packed).slice(4)));


// eval(function (p, a, c, k, e, d) {
//     e = function (c) {
//         return c.toString(36)
//     };
//     if (!''.replace(/^/, String)) {
//         while (c--) {
//             d[e(c)] = k[c] || e(c)
//         }
//         k = [function (e) {
//             return d[e]
//         }];
//         e = function () {
//             return '\\w+'
//         };
//         c = 1
//     }
//     ;
//     while (c--) {
//         if (k[c]) {
//             p = p.replace(new RegExp('\\b' + e(c) + '\\b', 'g'), k[c])
//         }
//     }
//     return p
// }('c[1]="e/b/a/d/g/m.f/0";c[2]="e/b/a/d/g/o.f/0";c[3]="e/b/a/d/g/l.f/0";c[4]="e/b/a/d/g/p.f/0";c[5]="e/b/a/d/g/k.f/0";c[6]="e/b/a/d/g/h.f/0";c[7]="e/b/a/d/g/i.f/0";c[8]="e/b/a/d/g/j.f/0";c[9]="e/b/a/d/g/n.f/0";c[q]="e/b/a/d/g/y.f/0";c[x]="e/b/a/d/g/z.f/0";c[v]="e/b/a/d/g/w.f/0";c[r]="e/b/a/d/g/s.f/0";c[t]="e/b/a/d/g/u.f/0";', 36, 36, '||||||||||06|2022|photosr|08|images2|jpg|09|4410ab9b42|44dc774ea8|44ee3bf26f|4408cb8969|4423ac5f28|4469b95f5d|44df0661d2|44540dd2a8|44aea07b90|10|13|45c4a59947|14|45e4caa3a1|12|456f1b902b|11|449ad225af|44be151ecf'.split('|'), 0, {}))
