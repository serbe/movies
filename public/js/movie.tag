<movie>
    <div each="{movie in movies}" class="movie border shadow my1" no-reorder>
        <div class="h2 m1">{movie.Name} (<a href="/search?year={movie.Year}">{movie.Year}</a>)</div>
        <div class="clearfix">
            <div class="justify desc sm-col sm-col-8 px1">
                <img src="/public/img/{movie.Poster}" alt="{movie.Name}">{movie.Description}
            </div>
            <div class="addon sm-col sm-col-4 px1">
                <div class="genre">Жанр:
                    <span each="{genre, i in movie.Genre}">
                        <a href="/search?genre={genre}">{genre}</a><span if={i < movie.Genre.length-1}>{', '}</span>
                    </span>
                </div>
                <div>Время: {movie.Duration}</div>
                <div if={movie.NNM != 0}>Рейтинг NNM: <span class="rating">{movie.NNM}</span></div>
                <div if={movie.Kinopoisk != 0}>Рейтинг Кинопоиска: {movie.Kinopoisk}</div>
                <div if={movie.IMDb != 0}>Рейтинг IMDb: {movie.IMDb}</div>
            </div>
        </div>
        <div class="pointer mx2 my1" onclick="switching(this)">
            <span>+ Показать дополнительную информацию</span>
        </div>
        <div class="toggle hide mb2 ml2 mr2">
            <div class="clearfix" if={movie.Country.length > 0}>
                <div class="sm-col sm-col-2 px1">Страна:</div>
                <div class="justify sm-col sm-col-10 px1">
                    <span each="{country, i in movie.Country}">
                        <a href="/search?country={country}">{country}</a><span if={i < movie.Country.length-1}>{', '}</span>
                    </span>
                </div>
            </div>
            <div class="clearfix" if={movie.Director.length > 0}>
                <div class="sm-col sm-col-2 px1">Режиссер:</div>
                <div class="justify sm-col sm-col-10 px1">
                    <span each="{director, i in movie.Director}">
                        <a href="/search?director={director}">{director}</a><span if={i < movie.Director.length-1}>{', '}</span>
                    </span>
                </div>
            </div>
            <div class="clearfix" if={movie.Producer.length > 0}>
                <div class="sm-col sm-col-2 px1">Продюссер:</div>
                <div class="justify sm-col sm-col-10 px1">
                    <span each="{producer, i in movie.Producer}">
                        <a href="/search?producer={producer}">{producer}</a><span if={i < movie.Producer.length-1}>{', '}</span>
                    </span>
                </div>
            </div>
            <div class="clearfix" if={movie.Actor.length > 0}>
                <div class="sm-col sm-col-2 px1">Актеры:</div>
                <div class="justify sm-col sm-col-10 px1">
                    <span each="{actor, i in movie.Actor}">
                        <a href="/search?actor={actor}">{actor}</a><span if={i < movie.Actor.length-1}>{', '}</span>
                    </span>
                </div>
            </div>
            <div class="clearfix my1">
                <div class="bold center sm-col sm-col-1 px1">Форум</div>
                <div class="bold center sm-col sm-col-1 px1">Magnet</div>
                <div class="bold center sm-col sm-col-1 px1">Размер</div>
                <div class="bold center sm-col sm-col-1 px1">Сиды</div>
                <div class="bold center sm-col sm-col-1 px1">Пиры</div>
                <div class="bold center sm-col sm-col-7 px1">Перевод</div>
            </div>
            <div class="clearfix" each="{torrent in movie.Torrent}">
                <div class="center sm-col sm-col-1">
                    <a href="http://nnmclub.to/forum/viewtopic.php?t={torrent.Href}">
                        <i class="fa fa-external-link"></i>
                    </a>
                </div>
                <div class="center sm-col sm-col-1">
                    <a href="magnet:?xt=urn:btih:{torrent.Magnet}">
                        <i class="fa fa-magnet"></i>
                    </a>
                </div>
                <div class="center sm-col sm-col-1">{torrent.Size}Мб</div>
                <div class="center sm-col sm-col-1">
                    <i class="fa fa-arrow-up"></i> {torrent.Seeders}
                </div>
                <div class="center sm-col sm-col-1">
                    <i class="fa fa-arrow-down"></i> {torrent.Leechers}
                </div>
                <div class="sm-col sm-col-7">{torrent.Translation}</div>
            </div>
        </div>
    </div>

    <script>
        scrollEnd(e) {
            console.log(e)
        }

        document.addEventListener(
            'scroll',
            function(ev)
            {
                if ((window.innerHeight + window.scrollY) >= document.body.scrollHeight && self.mcount > self.moffset)
                {
                    var url = "/movies?offset=" + self.moffset + "&limit=5";
                    var xmlhttp = new XMLHttpRequest();
                    xmlhttp.onreadystatechange = function() {
                        if (xmlhttp.readyState == 4 && xmlhttp.status == 200) {
                            var data = JSON.parse(xmlhttp.responseText);
                            if (data.Movies != null) {
                                for (var i = 0; i < data.Movies.length; i++) {
                                    self.movies.push(data.Movies[i])
                                }
                                self.mlimit = data.Limit
                                self.moffset = data.Offset
                                self.mcount = data.Count
                            riot.update()
                            }
                        }
                    }
                    xmlhttp.open("GET", url, true);
                    xmlhttp.send();
                }
            }.bind(this)
        )
        self.mlimit = opts.data.Limit
        self.moffset = opts.data.Offset
        self.mcount = opts.data.Count
        self.movies = opts.data.Movies
    </script>

</movie>