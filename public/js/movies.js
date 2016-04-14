

// {{define "content"}}
// <div class="uk-container uk-container-center uk-margin-top uk-margin-large-bottom">


//     <div class="uk-margin-large-bottom">
//         {{range $movie := .Movies}}
//         <div class="movie">
//             <hr class="movie-divider uk-margin-small-top uk-margin-small-bottom">
//             <h2 class="uk-text-large uk-margin-small-top uk-margin-small-bottom">{{$movie.Name}}{{if ne $movie.EngName ""}} / {{$movie.EngName}}{{end}}</h2>
//             <div class="uk-grid">
//                 <div class="uk-text-justify desc uk-width-7-10"><img src="/public/img/{{$movie.Poster}}" alt="{{$movie.Name}}">{{$movie.Description}}</div>
//                 <div class="addon uk-width-3-10">
//                     <div class="year uk-width-1-1">Год: <a href="/search?year={{$movie.Year}}">{{$movie.Year}}</a></div>
//                     <div class="genre uk-width-1-1">Жанр: {{range $index, $genre := $movie.Genre}}
//                         <a href="/search?genre={{$genre}}">{{$genre}}</a>{{if nolast $index $movie.Genre}},{{end}} {{end}}
//                     </div>
//                     <div class="uk-width-1-1">Продолжительность: {{$movie.Duration}}</div>
//                     {{if ne $movie.NNM 0.0}}
//                     <div class="uk-width-1-1">Рейтинг NNM: <span class="rating">{{$movie.NNM}}</span></div>
//                     {{end}} {{if ne $movie.Kinopoisk 0.0}}
//                     <div class="uk-width-1-1">Рейтинг Кинопоиска: {{$movie.Kinopoisk}}</div>
//                     {{end}} {{if ne $movie.IMDb 0.0}}
//                     <div class="uk-width-1-1">Рейтинг IMDb: {{$movie.IMDb}}</div>
//                     {{end}}
//                 </div>
//             </div>
//             <hr class="uk-margin-small-top uk-margin-small-bottom uk-width-1-1">
//             <div class="pointer uk-margin-large-left uk-margin-top-remove uk-margin-bottom-remove" onclick="switching(this)" data-uk-toggle="{target:'#i{{$movie.ID}}'}">
//                 + Показать дополнительную информацию
//             </div>
//             <div id="i{{$movie.ID}}" class="uk-hidden">
//                 <hr class="uk-width-1-1 uk-margin-small-top">
//                 <div class="uk-grid uk-margin-top">
//                     <div class="uk-width-1-10">Страна:</div>
//                     <div class="uk-text-justify uk-width-9-10">{{$movie.Country}}</div>
//                 </div>
//                 <div class="uk-grid uk-margin-top-remove">
//                     <div class="uk-width-1-10">Режиссер:</div>
//                     <div class="uk-text-justify uk-width-9-10">
//                         {{range $index, $director := $movie.Director}}
//                         <a href="/search?director={{$director}}">{{$director}}</a>{{if nolast $index $movie.Director}},{{end}}
//                         {{end}}
//                     </div>
//                 </div>
//                 <div class="uk-grid uk-margin-top-remove">
//                     <div class="uk-width-1-10">Актеры:</div>
//                     <div class="uk-text-justify uk-width-9-10">
//                         {{range $index, $actor := $movie.Actors}}
//                         <a href="/search?actor={{$actor}}">{{$actor}}</a>{{if nolast $index $movie.Actors}},{{end}} {{end}}
//                     </div>
//                 </div>
//                 <div class="uk-grid uk-margin-top-remove">
//                     <div class="uk-align-center uk-width-9-10 uk-margin-top">
//                         <table class="uk-table uk-table-condensed">
//                             <thead class="uk-text-small">
//                                 <tr>
//                                     <th class="uk-text-center uk-width-1-10">Форум</th>
//                                     <th class="uk-text-center uk-width-1-10">Magnet</th>
//                                     <th class="uk-text-center uk-width-1-10">Размер</th>
//                                     <th class="uk-text-center uk-width-1-10">Сиды</th>
//                                     <th class="uk-text-center uk-width-1-10">Пиры</th>
//                                     <th class="uk-text-center uk-width-4-10">Перевод</th>
//                                 </tr>
//                             </thead>
//                             <tbody>
//                                 {{range $index, $torrent := $movie.Torrents}}
//                                 <tr>
//                                     <td class="uk-table-middle uk-text-center">
//                                         <a class="uk-icon-external-link" href="http://nnmclub.to/forum/viewtopic.php?t={{$torrent.Href}}"></a>
//                                     </td>
//                                     <td class="uk-table-middle uk-text-center">
//                                         <a class="uk-icon-magnet" href="magnet:?xt=urn:btih:{{$torrent.Magnet}}"></a>
//                                     </td>
//                                     <td class="uk-table-middle uk-text-center">{{$torrent.Size}} Мб</td>
//                                     <td class="uk-table-middle uk-text-center">
//                                         <a class="uk-icon-long-arrow-up"></a> {{$torrent.Seeders}}
//                                     </td>
//                                     <td class="uk-table-middle uk-text-center">
//                                         <a class="uk-icon-long-arrow-down"></a> {{$torrent.Leechers}}
//                                     </td>
//                                     <td class="uk-table-middle">{{$torrent.Translation}}</td>
//                                 </tr>
//                                 {{end}}
//                             </tbody>
//                         </table>
//                     </div>
//                 </div>
//             </div>
//         </div>
//         {{end}}
//         <hr class="movie-divider uk-margin-top uk-margin-bottom">
//     </div>
//     <!--<ul class="uk-pagination" data-uk-pagination="{items:100, itemsOnPage:50, currentPage:1}">
//         <li><a href="">...</a></li>
//         <li class="uk-active"><span>...</span></li>
//         <li class="uk-disabled"><span>...</span></li>
//         <li><span>...</span></li>
//     </ul>-->
// </div>
// {{end}}

"use strict";

var NewComponent = React.createClass({
  displayName: "NewComponent",

  render: function render() {
    return React.createElement(
      "div",
      { className: "movie" },
      React.createElement("hr", { className: "movie-divider uk-margin-small-top uk-margin-small-bottom" }),
      React.createElement(
        "h2",
        { className: "uk-text-large uk-margin-small-top uk-margin-small-bottom" },
        "movie.Name movie.EngName \"\" / movie.EngName"
      ),
      React.createElement(
        "div",
        { className: "uk-grid" },
        React.createElement(
          "div",
          { className: "uk-text-justify desc uk-width-7-10" },
          React.createElement("img", { src: "/public/img/movie.Poster", alt: "movie.Name" }),
          "movie.Description"
        ),
        React.createElement(
          "div",
          { className: "addon uk-width-3-10" },
          React.createElement(
            "div",
            { className: "year uk-width-1-1" },
            "Год: ",
            React.createElement(
              "a",
              { href: "/search?year={{$movie.Year}}" },
              "movie.Year"
            )
          ),
          React.createElement(
            "div",
            { className: "genre uk-width-1-1" },
            "Жанр: movie.Genre",
            React.createElement(
              "a",
              { href: "/search?genre={{$genre}}" },
              "genre"
            )
          ),
          React.createElement(
            "div",
            { className: "uk-width-1-1" },
            "Продолжительность: movie.Duration"
          ),
          React.createElement(
            "div",
            { className: "uk-width-1-1" },
            "Рейтинг NNM: ",
            React.createElement(
              "span",
              { className: "rating" },
              "movie.NNM"
            )
          ),
          React.createElement(
            "div",
            { className: "uk-width-1-1" },
            "Рейтинг Кинопоиска: movie.Kinopoisk"
          ),
          React.createElement(
            "div",
            { className: "uk-width-1-1" },
            "Рейтинг IMDb: movie.IMDb"
          )
        )
      ),
      React.createElement("hr", { className: "uk-margin-small-top uk-margin-small-bottom uk-width-1-1" }),
      React.createElement(
        "div",
        { className: "pointer uk-margin-large-left uk-margin-top-remove uk-margin-bottom-remove", onclick: "switching(this)", "data-uk-toggle": "{target:'#i{{$movie.ID}}'}" },
        "+ Показать дополнительную информацию"
      ),
      React.createElement(
        "div",
        { id: "i{{$movie.ID}}", className: "uk-hidden" },
        React.createElement("hr", { className: "uk-width-1-1 uk-margin-small-top" }),
        React.createElement(
          "div",
          { className: "uk-grid uk-margin-top" },
          React.createElement(
            "div",
            { className: "uk-width-1-10" },
            "Страна:"
          ),
          React.createElement(
            "div",
            { className: "uk-text-justify uk-width-9-10" },
            "movie.Country"
          )
        ),
        React.createElement(
          "div",
          { className: "uk-grid uk-margin-top-remove" },
          React.createElement(
            "div",
            { className: "uk-width-1-10" },
            "Режиссер:"
          ),
          React.createElement(
            "div",
            { className: "uk-text-justify uk-width-9-10" },
            React.createElement(
              "a",
              { href: "/search?director={{$director}}" },
              "director"
            )
          )
        ),
        React.createElement(
          "div",
          { className: "uk-grid uk-margin-top-remove" },
          React.createElement(
            "div",
            { className: "uk-width-1-10" },
            "Актеры:"
          ),
          React.createElement(
            "div",
            { className: "uk-text-justify uk-width-9-10" },
            React.createElement(
              "a",
              { href: "/search?actor={{$actor}}" },
              "actor"
            )
          )
        ),
        React.createElement(
          "div",
          { className: "uk-grid uk-margin-top-remove" },
          React.createElement(
            "div",
            { className: "uk-align-center uk-width-9-10 uk-margin-top" },
            React.createElement(
              "table",
              { className: "uk-table uk-table-condensed" },
              React.createElement(
                "thead",
                { className: "uk-text-small" },
                React.createElement(
                  "tr",
                  null,
                  React.createElement(
                    "th",
                    { className: "uk-text-center uk-width-1-10" },
                    "Форум"
                  ),
                  React.createElement(
                    "th",
                    { className: "uk-text-center uk-width-1-10" },
                    "Magnet"
                  ),
                  React.createElement(
                    "th",
                    { className: "uk-text-center uk-width-1-10" },
                    "Размер"
                  ),
                  React.createElement(
                    "th",
                    { className: "uk-text-center uk-width-1-10" },
                    "Сиды"
                  ),
                  React.createElement(
                    "th",
                    { className: "uk-text-center uk-width-1-10" },
                    "Пиры"
                  ),
                  React.createElement(
                    "th",
                    { className: "uk-text-center uk-width-4-10" },
                    "Перевод"
                  )
                )
              ),
              React.createElement(
                "tbody",
                null,
                React.createElement(
                  "tr",
                  null,
                  React.createElement(
                    "td",
                    { className: "uk-table-middle uk-text-center" },
                    React.createElement("a", { className: "uk-icon-external-link", href: "http://nnmclub.to/forum/viewtopic.php?t={{$torrent.Href}}" })
                  ),
                  React.createElement(
                    "td",
                    { className: "uk-table-middle uk-text-center" },
                    React.createElement("a", { className: "uk-icon-magnet", href: "magnet:?xt=urn:btih:{{$torrent.Magnet}}" })
                  ),
                  React.createElement(
                    "td",
                    { className: "uk-table-middle uk-text-center" },
                    "torrent.Size Мб"
                  ),
                  React.createElement(
                    "td",
                    { className: "uk-table-middle uk-text-center" },
                    React.createElement("a", { className: "uk-icon-long-arrow-up" }),
                    " torrent.Seeders"
                  ),
                  React.createElement(
                    "td",
                    { className: "uk-table-middle uk-text-center" },
                    React.createElement("a", { className: "uk-icon-long-arrow-down" }),
                    " torrent.Leechers"
                  ),
                  React.createElement(
                    "td",
                    { className: "uk-table-middle" },
                    "torrent.Translation"
                  )
                )
              )
            )
          )
        )
      )
    );
  }
});