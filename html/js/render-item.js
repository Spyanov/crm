Vue.component('renderitem',{
    props:{
        id:Number,
        client: String,
        title: String,
        desc: String,
        price: Number,
        start: String,
        end: String,
        statusProp: String,
        resultProp: String,
        statusArr: Array,
        resultArr: Array,
    },
    data:function(){
        return{
            currentId:null,
            currentClient:null,
            currentTitle:null,
            currentDesc:null,
            currentPrice:null,
            currentStart:null,
            currentEnd:null,
            currentStatus:null,
            currentResult:null,
        }
    },

    methods:{
        sendUpadte:function n(){
            $.ajax({
                url : "/update",
                type: "POST",

                timeout:5000,

                data:`{
                  "id":"`+this.currentId +`",
                  "client":"`+this.currentClient+`",
                  "dealTitle":"`+this.currentTitle+`",
                  "dealDesc":"`+this.currentDesc+`",
                  "price":`+this.currentPrice+`,
                  "status":"`+this.currentStatus+`",
                  "result":"`+this.currentResult+`"}`,

                contentType: false,

                cache: false,
                processData:false,


            }).done(function(data, textStatus, jqXHR){ //
                if (jqXHR.status == 201){
                    $('.modal').modal('hide')
                    app.loading();


                } else {
                    alert("Что-то пошло не так: ")
                }

            }).fail(function(data, textStatus, jqXHR) {
                alert("Не удалось записать: ");
            })
        }
    },

    template:`
          <div  class="item">
            <div>
              <a data-toggle="modal" :data-target="'#modal'+this.id" class="color">
                <div class="item-body" :id="'id'+this.id">

                  <div class="item-client-name">{{this.client}}</div>

                  <div class="item-title">{{this.title}}</div>

                  <div class="item-desc">{{this.desc}}</div>


                  <div class="item-price"><i class="fas fa-coins"></i> {{this.price}}</div>


                </div>
               </a>
            </div>



              <!-- Modal -->
              <div class="modal fade" :id="'modal'+this.id" tabindex="-1" role="dialog" :aria-labelledby="'label'+id" aria-hidden="true">
                <div class="modal-dialog modal-lg" role="document">
                    <div class="modal-content">



              <form action="/" method="POST" :id="'updateItem'+id" class="text-center border border-light p-5">
              <div class="container inputBlock">
              <div class="row">
              <div class="col-12" >
              </div>
                      <input name="id" :value="id" class="hide">
              </div>
              <div class="row">
              <div class="col-12">
              <input type="text" class="form-control mb-4" :placeholder="client" name="client" v-model:value="currentClient">

              </div>
              </div>
              <div class="row">
              <div class="col-12">
              <input type="text" class="form-control mb-4" :placeholder="title" name="title" v-model:value="currentTitle">

              </div>
              </div>
              <div class="row">
              <div class="col-12">
              <input type="text" class="form-control mb-4" :placeholder="desc" name="desc" v-model:value="desc">

              </div>
              </div>
              <div class="row">
              <div class="col-12">
              <input type="number" class="form-control mb-4" :placeholder="price" name="price" v-model:value="price" maxlength="7">

              </div>
              </div>
              <div class="row">
              <div class="col-12">
              <input type="text" class="form-control mb-4" placeholder="Начало" name="start" >

              </div>
              </div>
              <div class="row">
              <div class="col-12">
              <input type="text" class="form-control mb-4" placeholder="Конец" name="end">

              </div>
              </div>
              <div class="row">
              <div class="col-12">
              <select class="browser-default custom-select mb-4" v-model="statusProp" name="status">
              <option value="" disabled="">Результат</option>
              <option v-for="statusList in this.statusArr" :value="statusList" >{{statusList}}</option>
              </select>

              </div>
              </div>
              <div class="row">
              <div class="col-12">
              <select class="browser-default custom-select mb-4" v-model="resultProp" name="result">
              <option value="" disabled="">Стадия</option>
              <option v-for="resultList in this.resultArr" :value="resultList" >{{resultList}}</option>
              </select>
              </div>
              </div>
              </div>

              <button type="button" class="btn btn-secondary btn-sm" data-dismiss="modal"><i class="far fa-window-close"></i></button>
              <button type="button" class="btn btn-primary btn-sm" v-on:click="this.sendUpdate"><i class="far fa-save"></i></button>
              </form>


                    </div>
                  </div>
                </div>
              </div>
              <!-- .Modal -->


              `
});