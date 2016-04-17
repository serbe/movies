<widget-panel>
    <div class="panel panel-{ type }" onclick="{ click }">
      <div class="panel-heading">
        <h3 class="panel-title">{ opts.title }</h3>
      </div>
      <div class="panel-body">{ opts.content }</div>
    </div>

    this.type = opts.type || 'default'

    click(e){
        this.type = 'success';
    }
</widget-panel>