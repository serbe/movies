<movie>
    <div each="{movie in movies}" class="movie" no-reorder>
        <hr class="movie-divider uk-margin-small-top uk-margin-small-bottom">
        <h2 class="uk-text-large uk-margin-small-top uk-margin-small-bottom">{movie.Name}</h2>
        <div class="columns">
            <div class="uk-text-justify desc column is-9"><img src="/public/img/{movie.Poster}" alt="{movie.Name}">{movie.Description}</div>
            <div class="addon column is-3">
                <div class="year">Год: <a href="/search?year={movie.Year}">{movie.Year}</a></div>
                <div class="genre">Жанр:
                    <span each="{genre, i in movie.Genre}">
                        <a href="/search?genre={genre}">{genre}</a><span if={i < movie.Genre.length-1}>{', '}</span>
                    </span>
                </div>
                <div class="">Продолжительность: {movie.Duration}</div>
                <div if={movie.NNM != 0} class="">Рейтинг NNM: <span class="rating">{movie.NNM}</span></div>
                <div if={movie.Kinopoisk != 0} class="">Рейтинг Кинопоиска: {movie.Kinopoisk}</div>
                <div if={movie.IMDb != 0} class="">Рейтинг IMDb: {movie.IMDb}</div>
            </div>
        </div>
        <hr class="uk-margin-small-top uk-margin-small-bottom">
        <div class="pointer uk-margin-large-left uk-margin-top-remove uk-margin-bottom-remove" onclick="switching(this)" data-uk-toggle="\{target:'#i{movie.ID}'\}">
            <span>+ Показать дополнительную информацию</span>
        </div>
        <div id="i{movie.ID}" class="uk-hidden">
            <hr class="uk-margin-small-top">
            <div class="columns uk-margin-top-remove" if={movie.Country.length > 0}>
                <div class="column is-2">Страна:</div>
                <div class="uk-text-justify column is-10">
                    <span each="{country, i in movie.Country}">
                        <a href="/search?country={country}">{country}</a><span if={i < movie.Country.length-1}>{', '}</span>
                    </span>
                </div>
            </div>
            <div class="columns uk-margin-top-remove" if={movie.Director.length > 0}>
                <div class="column is-2">Режиссер:</div>
                <div class="uk-text-justify column is-10">
                    <span each="{director, i in movie.Director}">
                        <a href="/search?director={director}">{director}</a><span if={i < movie.Director.length-1}>{', '}</span>
                    </span>
                </div>
            </div>
            <div class="columns uk-margin-top-remove" if={movie.Producer.length > 0}>
                <div class="column is-2">Продюссер:</div>
                <div class="uk-text-justify column is-10">
                    <span each="{producer, i in movie.Producer}">
                        <a href="/search?producer={producer}">{producer}</a><span if={i < movie.Producer.length-1}>{', '}</span>
                    </span>
                </div>
            </div>
            <div class="columns uk-margin-top-remove" if={movie.Actor.length > 0}>
                <div class="column is-2">Актеры:</div>
                <div class="uk-text-justify column is-10">
                    <span each="{actor, i in movie.Actor}">
                        <a href="/search?actor={actor}">{actor}</a><span if={i < movie.Actor.length-1}>{', '}</span>
                    </span>
                </div>
            </div>
            <div class="columns uk-margin-top-remove">
                <div class="uk-align-center column is-10 uk-margin-top">
                    <table class="uk-table uk-table-condensed">
                        <thead class="uk-text-small">
                            <tr>
                                <th class="uk-text-center uk-width-1-10">Форум</th>
                                <th class="uk-text-center uk-width-1-10">Magnet</th>
                                <th class="uk-text-center uk-width-1-10">Размер</th>
                                <th class="uk-text-center uk-width-1-10">Сиды</th>
                                <th class="uk-text-center uk-width-1-10">Пиры</th>
                                <th class="uk-text-center uk-width-4-10">Перевод</th>
                            </tr>
                        </thead>
                        <tbody>
                            <tr each="{torrent in movie.Torrent}">
                                <td class="uk-table-middle uk-text-center">
                                    <a class="uk-icon-external-link" href="http://nnmclub.to/forum/viewtopic.php?t={torrent.Href}"></a>
                                </td>
                                <td class="uk-table-middle uk-text-center">
                                    <a class="uk-icon-magnet" href="magnet:?xt=urn:btih:{torrent.Magnet}"></a>
                                </td>
                                <td class="uk-table-middle uk-text-center">{torrent.Size}Мб</td>
                                <td class="uk-table-middle uk-text-center">
                                    <a class="uk-icon-long-arrow-up"></a> {torrent.Seeders}
                                </td>
                                <td class="uk-table-middle uk-text-center">
                                    <a class="uk-icon-long-arrow-down"></a> {torrent.Leechers}
                                </td>
                                <td class="uk-table-middle">{torrent.Translation}</td>
                            </tr>
                        </tbody>
                    </table>
                </div>
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