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
            currentId:this.id,
            currentClient:this.client,
            currentTitle:this.title,
            currentDesc:this.desc,
            currentPrice:this.price,
            currentStart:null,
            currentEnd:null,
            currentStatus:this.statusProp,
            currentResult:this.resultProp,
        }
    },

    methods:{
        sendUpdate:function(event){
            $.ajax({
                url : "/update",
                type: "POST",

                timeout:5000,

                data:`{
                  "id":`+this.currentId +`,
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
                if (jqXHR.status == 200){
                    $('.modal').modal('hide')
                    app.loading();
                } else {
                    alert("Что-то пошло не так: ")
                }

            }).fail(function(data, textStatus, jqXHR) {
                alert("Не удалось записать: ");
            })
        },
        delItem:function () {
            $.ajax({
                url : "/del",
                type: "POST",

                timeout:5000,

                data:`{"id":`+this.currentId +`}`,

                contentType: false,

                cache: false,
                processData:false,


            }).done(function(data, textStatus, jqXHR){ //
                if (jqXHR.status == 200){
                    $('.modal').modal('hide')
                    console.log("запись" + this.currentId +" удалена");
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
               <!-- item  -->
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
              <!-- item  -->
              <!-- Modal item -->
              <div class="modal fade" :id="'modal'+this.id" tabindex="-1" role="dialog" :aria-labelledby="'label'+id" aria-hidden="true">
                 <div class="modal-dialog modal-lg" role="document">
                    <div class="modal-content">
                    
                        <div class="d-flex flex-row-reverse">
                            <button type="button" class="btn btn-danger btn-sm" @click="this.delItem">DELETE</button>
                        </div>
                        
                      <form action="/" method="POST" :id="'updateItem'+id" class="text-center border border-light p-5">
                        <div class="container inputBlock">
                          <div class="row">
                          <div class="col-12 " >
                              </div>
                                      <input name="id" v-bind:value="currentId" class="hide">
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
                          <input type="text" class="form-control mb-4" :placeholder="desc" name="desc" v-model:value="currentDesc">
            
                          </div>
                          </div>
                          <div class="row">
                          <div class="col-12">
                          <input type="number" class="form-control mb-4" :placeholder="price" name="price" v-model:value="currentPrice" maxlength="7">
            
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
                          <select class="browser-default custom-select mb-4" v-model="currentStatus" name="status">
                          <option value="" disabled="">Результат</option>
                          <option v-for="statusList in this.statusArr" :value="statusList" >{{statusList}}</option>
                          </select>
            
                          </div>
                          </div>
                          <div class="row">
                          <div class="col-12">
                          <select class="browser-default custom-select mb-4" v-model="currentResult" name="result">
                          <option value="" disabled="">Стадия</option>
                          <option v-for="resultList in this.resultArr" :value="resultList" >{{resultList}}</option>
                          </select>
                          </div>
                          </div>
                          </div>
            
                        <div class="row">
                        <div class="col-6"></div>
                        <div class="col-6">
                            <button type="button" class="btn btn-light btn-lg" data-dismiss="modal">Отмена</button>
                            <button type="button" class="btn btn-primary btn-lg" @click="this.sendUpdate">Сохранить</button></div>
                        </div>
                        <div class="row">

                        </div>
                        
                        
                        <br>
                        
                      </form>
                    </div>
                 </div>
              </div>
               <!-- .Modal item -->
          </div>
             


              `
});