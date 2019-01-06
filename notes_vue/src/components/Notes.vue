<template>
  <div>
    <div class="container">
      <form v-on:submit.prevent="noop">
        <input class="new_note_textbox" type="text" name="new_note" v-on:keyup.enter="addNote(newnote)" v-model="newnote">
      </form>

      <div v-for="note in notes" class="card">
        <p class="card_text">{{note.text}}</p>
      </div>

    </div>


  </div>
</template>

<script>
  export default {
    name: 'Notes',
    data () {
      return {
        newnote: "",
        notes: [],
      }
    },
    mounted(){
      var data = null;
      var self = this;

      var xhr = new XMLHttpRequest();
      xhr.withCredentials = true;

      xhr.addEventListener("readystatechange", function () {
        if (this.readyState === 4) {
          self.notes = JSON.parse(this.response).data;

          self.notes.sort(function (a, b) {
            return b.id - a.id;
          })
        }
      });

      xhr.open("GET", "http://localhost:3000/api/v1/notes/");

      xhr.send(data);
    },
    methods: {
      addNote: function (note) {
        var data = "text="+note;
        var self = this;
        var xhr = new XMLHttpRequest();
        xhr.withCredentials = true;

        xhr.addEventListener("readystatechange", function () {
          if (this.readyState === 4) {
            console.log(this.responseText);
            self.updateNotes();
            self.newnote = "";
          }
        });

        xhr.open("POST", "http://localhost:3000/api/v1/notes/");
        xhr.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");

        xhr.send(data);


      },

      noop: function () {

      },

      updateNotes: function () {
        var data = null;
        var self = this;

        var xhr = new XMLHttpRequest();
        xhr.withCredentials = true;

        xhr.addEventListener("readystatechange", function () {
          if (this.readyState === 4) {
            self.notes = JSON.parse(this.response).data;
            self.notes.sort(function (a, b) {
              return b.id - a.id;
            })
          }
        });

        xhr.open("GET", "http://localhost:3000/api/v1/notes/");

        xhr.send(data);
      },
    }
  }
</script>


<style scoped>
  .new_note_textbox{
    width: 100%;
    padding: 20px;
    margin-bottom: 30px;
    box-sizing: border-box;

  }

  .container{
    padding-left: 20%;
    padding-right: 20%;
  }

  .card {
    width: 100%;
    border: 1px solid #e8e8e8;
    border-radius: 5px; /* 5px rounded corners */
    background-color: #f8f8f8;
    margin-bottom: 15px;
  }

  .card_text {
    padding: 10px;
    margin: 20px;
    text-align: left;
    color: #1d1d1d;
  }

</style>
