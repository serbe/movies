<movie>
    <div each="{ movie in movies }" class="movie">
        <hr class="movie-divider uk-margin-small-top uk-margin-small-bottom">
        <h2 class="uk-text-large uk-margin-small-top uk-margin-small-bottom">{ movie.Name }</h2>
        <div class="uk-grid">
            <div class="uk-text-justify desc uk-width-7-10"><img src="/public/img/{ movie.Poster }" alt="{ movie.Name }">{ movie.Description }</div>
            <div class="addon uk-width-3-10">
                <div class="year uk-width-1-1">Год: <a href="/search?year={ movie.Year }">{ movie.Year }</a></div>
                <div class="genre uk-width-1-1">Жанр: { movie.Genre }
                    <a each="{ genre in movie.Genre }" href="/search?genre={ genre }">{ genre }</a>
                </div>
                <div class="uk-width-1-1">Продолжительность: { movie.Duration }</div>
                <div if={ movie.NNM != 0 } class="uk-width-1-1">Рейтинг NNM: <span class="rating">{ movie.NNM }</span></div>
                <div if={ movie.Kinopoisk != 0 } class="uk-width-1-1">Рейтинг Кинопоиска: { movie.Kinopoisk }</div>
                <div if={ movie.IMDb != 0 } class="uk-width-1-1">Рейтинг IMDb: { movie.IMDb }</div>
            </div>
        </div>
        <hr class="uk-margin-small-top uk-margin-small-bottom uk-width-1-1">
        <div class="uk-hidden">
            <hr class="uk-width-1-1 uk-margin-small-top">
            <div class="uk-grid uk-margin-top">
                <div class="uk-width-1-10">Страна:</div>
                <div class="uk-text-justify uk-width-9-10">{ movie.Country }</div>
            </div>
            <div class="uk-grid uk-margin-top-remove">
                <div class="uk-width-1-10">Режиссер:</div>
                <div class="uk-text-justify uk-width-9-10">
                    <a each="{ director in movie.Director }" href="/search?director={ director }">{ director }</a>
                </div>
            </div>
            <div class="uk-grid uk-margin-top-remove">
                <div class="uk-width-1-10">Актеры:</div>
                <div class="uk-text-justify uk-width-9-10">
                    <a each="{ actor in movie.Actors }" href="/search?actor={ actor }">{ actor }</a>
                </div>
            </div>
            <div class="uk-grid uk-margin-top-remove">
                <div class="uk-align-center uk-width-9-10 uk-margin-top">
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
                            <tr each="{ torrent in movie.Torrents }">
                                <td class="uk-table-middle uk-text-center">
                                    <a class="uk-icon-external-link" href="http://nnmclub.to/forum/viewtopic.php?t={ torrent.Href }"></a>
                                </td>
                                <td class="uk-table-middle uk-text-center">
                                    <a class="uk-icon-magnet" href="magnet:?xt=urn:btih:{ torrent.Magnet }"></a>
                                </td>
                                <td class="uk-table-middle uk-text-center">{ torrent.Size } Мб</td>
                                <td class="uk-table-middle uk-text-center">
                                    <a class="uk-icon-long-arrow-up"></a> { torrent.Seeders }
                                </td>
                                <td class="uk-table-middle uk-text-center">
                                    <a class="uk-icon-long-arrow-down"></a> { torrent.Leechers }
                                </td>
                                <td class="uk-table-middle">{ torrent.Translation }</td>
                            </tr>
                        </tbody>
                    </table>
                </div>
            </div>
        </div>
    </div>
    
    <script>
        self.movies = opts.data
        click (e) {
            alert('Hi!')
        }
    </script>
    
</movie>