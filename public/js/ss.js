'use strict';

function switching(e) {
    var movie = e.parentNode;
    var desc = movie.getElementsByClassName('desc');
    var addon = movie.getElementsByClassName('addon');
    var pointer = movie.getElementsByClassName('pointer');
    var dch = desc[0].clientHeight;
    var dsh = desc[0].scrollHeight;
    var ach = addon[0].clientHeight;
    var ash = addon[0].scrollHeight;
    if ((dsh == dch && dch != 81) || (ash == ach && ach != 81)) {
        pointer[0].innerHTML = '+ Показать дополнительную информацию'
        desc[0].style.height = '81px';
        addon[0].style.height = '81px';
    } else {
        pointer[0].innerHTML = '- Скрыть дополнительную информацию'
        if (dsh > ash) {
            desc[0].style.height = dsh + 'px';
            addon[0].style.height = dsh + 'px';
        } else {
            desc[0].style.height = ash + 'px';
            addon[0].style.height = ash + 'px';
        }
    }
}

function filter(phrase) {
    var movies = document.getElementsByClassName('movie');
    var words = phrase.value.toLowerCase().split(" ");
    var ele;
    for (var r = 0; r < movies.length; r++) {
        ele = movies[r].innerHTML.replace(/<[^>]+>/g, "");
        var displayStyle = 'none';
        for (var i = 0; i < words.length; i++) {
            if (ele.toLowerCase().indexOf(words[i]) >= 0 && movies[r].classList.contains('hy') === false && movies[r].classList.contains('hg') === false && movies[r].classList.contains('hr') === false)
                displayStyle = '';
            else {
                displayStyle = 'none';
                break;
            }
        }
        if (displayStyle == 'none' && movies[r].classList.contains('hs') === false) {
            movies[r].classList.add('hs')
        }
        if (displayStyle == '') {
            movies[r].classList.remove('hs')
        }
        movies[r].style.display = displayStyle;
    }
}

function filterColumnGenre(text, i) {
    var movies = document.getElementsByClassName('movie');
    var text = text.toLowerCase();
    var ele;
    for (var r = 0; r < movies.length; r++) {
        if (text != '') {
            ele = movies[r].getElementsByClassName('genre')[0].innerHTML.replace(/<[^>]+>/g, "");
            if (ele.toLowerCase().indexOf(text) >= 0 && movies[r].classList.contains('hy') === false && movies[r].classList.contains('hs') === false && movies[r].classList.contains('hr') === false) {
                movies[r].classList.remove('hg')
                movies[r].style.display = '';
            } else {
                movies[r].style.display = 'none';
                if (movies[r].classList.contains('hg') === false) {
                    movies[r].classList.add('hg')
                }
            }
        } else {
            movies[r].classList.remove('hg')
            if (movies[r].classList.contains('hy') === false && movies[r].classList.contains('hs') === false && movies[r].classList.contains('hr') === false) {
                movies[r].style.display = '';
            }
        }
    }
}

function filterColumnYear(text, i) {
    var movies = document.getElementsByClassName('movie');
    var ele;
    for (var r = 0; r < movies.length; r++) {
        if (text != '') {
            ele = movies[r].getElementsByClassName('year')[0].getElementsByTagName("a")[0].innerText
            if (ele == text && movies[r].classList.contains('hg') === false && movies[r].classList.contains('hs') === false && movies[r].classList.contains('hr') === false) {
                movies[r].classList.remove('hy')
                movies[r].style.display = '';
            } else {
                movies[r].style.display = 'none';
                if (movies[r].classList.contains('hy') === false) {
                    movies[r].classList.add('hy')
                }
            }
        } else {
            movies[r].classList.remove('hy')
            if (movies[r].classList.contains('hg') === false && movies[r].classList.contains('hs') === false && movies[r].classList.contains('hr') === false) {
                movies[r].style.display = '';
            }
        }
    }
}

function filterColumnRating(text, i) {
    var movies = document.getElementsByClassName('movie');
    var ele;
    var value = parseFloat(text.replace('&gt;', '').replace('>', ''));
    for (var r = 0; r < movies.length; r++) {
        if (value != 0) {
            ele = movies[r].getElementsByClassName('rating')[0].innerText.replace(/<[^>]+>/g, "");
            if (parseFloat(ele) >= value && movies[r].classList.contains('hg') === false && movies[r].classList.contains('hs') === false && movies[r].classList.contains('hy') === false) {
                movies[r].classList.remove('hr')
                movies[r].style.display = '';
            } else {
                movies[r].style.display = 'none';
                if (movies[r].classList.contains('hr') === false) {
                    movies[r].classList.add('hr')
                }
            }
        } else {
            movies[r].classList.remove('hr')
            if (movies[r].classList.contains('hg') === false && movies[r].classList.contains('hs') === false && movies[r].classList.contains('hy') === false) {
                movies[r].style.display = '';
            }
        }
    }
}

function changeGenre() {
    var e = document.getElementById("genre");
    var value = e.options[e.selectedIndex].value;
    var text = e.options[e.selectedIndex].text;
    if (value === '-1') {
        filterColumnGenre('', 5)
    } else {
        filterColumnGenre(text, 5)
    }
}

function changeYear() {
    var e = document.getElementById("year");
    var value = e.options[e.selectedIndex].value;
    var text = e.options[e.selectedIndex].text;
    if (value === '-1') {
        filterColumnYear('', 2)
    } else {
        filterColumnYear(text, 2)
    }
}

function changeRating() {
    var e = document.getElementById("rating");
    var value = e.options[e.selectedIndex].value;
    var text = e.options[e.selectedIndex].text;
    if (value === '0') {
        filterColumnRating('0', 3)
    } else {
        filterColumnRating(text, 3)
    }
}
