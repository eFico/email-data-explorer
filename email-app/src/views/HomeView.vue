<template>
  <main class="mx-8 text-white">
    <div class="pt-4 mb-8 relative">
      <input type="text" v-model="searchQuery" @keyup.enter="getSearchResults" placeholder="Search a email"
        class="py-2 px-1 w-full bg-transparent border-b focus:border-weather-secondary focus:outline-none focus:shadow-[0px_1px_0_0_#004E71]" />
    </div>

    <div class="w-full gap-4 flex">
      <div class="w-3/5 h-full ">
        <div class="w-full" v-if="isLoaded">
          <table class="table-fixed border-spacing-y-3 w-full text-gray-400 border-separate space-y-6 text-sm">
            <thead class="bg-gray-800 text-gray-500">
              <tr>
                <th class="p-3 w-1/4 cursor-pointer" @click="sort('Subject')">Subject</th>
                <th class="p-3 text-left w-1/4 cursor-pointer" @click="sort('From')">From</th>
                <th class="p-3 text-left w-1/2 cursor-pointer" @click="sort('To')">To</th>

              </tr>
            </thead>
            <tbody>
              <tr v-for="(row, index) in dataTable.records" :key="index" @click="selectRow(row, index)"
                v-bind:class="{ 'bg-gray-800': index !== seleccionada, 'bg-weather-secondary': index === seleccionada }">
                <td class="p-3 truncate">
                  <div class="flex align-items-center">
                    <img class="rounded-full h-12 w-12  object-cover"
                      src="https://icon-library.com/images/email-icon-images/email-icon-images-4.jpg"
                      alt="unsplash image">
                    <div class="ml-3">
                      <div class="truncate">{{ row._source.Subject }}</div>
                      <div class="text-gray-500 truncate">{{ row._source.Date }}</div>
                    </div>
                  </div>
                </td>
                <td class="p-3 truncate">{{ row._source.From }}</td>
                <td class="p-3 font-bold truncate">{{ row._source.To }}</td>

              </tr>

            </tbody>
          </table>

          <nav class="flex items-center justify-between py-4 ">
            <button
              class="px-4 py-2 font-bold text-white rounded-full bg-cyan-500 hover:bg-cyan-600 focus:outline-none focus:shadow-outline"
              @click="paginaAnterior" :disabled="paginaActual === 1">
              Anterior
            </button>
            <span class="px-4 py-2 font-bold text-white">Página {{ Math.ceil(paginaActual / elementosPorPagina) }} de {{ numPaginas }}</span>

            <button
              class="px-4 py-2 font-bold text-white rounded-full bg-cyan-500 hover:bg-cyan-600 focus:outline-none focus:shadow-outline"
              @click="paginaSiguiente" :disabled="paginaActual === numPaginas">
              Siguiente
            </button>
          </nav>

        </div>
      </div>
      <div class="w-2/5 divide-y divide-dashed ">
        <div class="my-2">
          {{ rowSelected? rowSelected.Subject : 'Subject' }}
        </div>
        <div class="my-2">
          {{ rowSelected? rowSelected.Date : 'Date' }}
        </div>
        <div class="my-2 overflow-y-scroll h-60 bg-scroll ">
          {{ rowSelected && rowSelected.body }}
        </div>

      </div>
    </div>

  </main>


</template>

<script setup>
import { ref, onMounted } from "vue";
import axios from "axios";

const searchQuery = ref("BRUCE");

const dataTable = ref(null);
const isLoaded = ref(null);

const rowSelected = ref(null);
const asDesc = ref(1);

const paginaActual = ref(1);
const elementosPorPagina = ref(10);
const totalRegistros = ref(null);
const numPaginas = ref(null);

const seleccionada = ref(null);
onMounted(() => {
  callAPI()
})

function selectRow(row, index) {
  console.log(index)
  rowSelected.value = row._source
  seleccionada.value = index
}
function sort(columna) {
  console.log('sort')
  console.log(dataTable.value)
  asDesc.value = -1 * asDesc.value
  dataTable.value.records.sort((a, b) => {
    if (a._source[columna] < b._source[columna]) return -1 * asDesc.value
    if (a._source[columna] > b._source[columna]) return 1 * asDesc.value
    return 0
  })
}


function paginaSiguiente() {
  seleccionada.value = null
  rowSelected.value = null
  paginaActual.value=paginaActual.value+elementosPorPagina.value
  callAPI()
}
function paginaAnterior() {
  seleccionada.value = null
  rowSelected.value = null
  paginaActual.value=paginaActual.value-elementosPorPagina.value
  callAPI()
}

const getSearchResults = () => {
  paginaActual.value = 1
  callAPI()
};

const callAPI = async () => {

  const config = {
    headers: {
      'Content-Type': 'application/json'
    }
  }

  // Crea el cuerpo de la solicitud
  const body = {
    page: paginaActual.value-1,
    Size: elementosPorPagina.value,
    Query: searchQuery.value
  }
  await axios.post('http://localhost:8080/emails', body, config)
    .then(response => {
      // La llamada se ha completado correctamente y se ha recibido una respuesta del servidor
      dataTable.value = response.data
      totalRegistros.value = response.data.total
      numPaginas.value = Math.ceil(totalRegistros.value / elementosPorPagina.value)

      isLoaded.value = true

    })
    .catch(error => {
      // Ha ocurrido un error al hacer la llamada
      console.log(error)
    })
};

</script>

<style lang="css" scoped>
.table {
  border-spacing: 0 10px;
}

i {
  font-size: 1rem !important;
}

.table tr {
  border-radius: 20px;
}

tr td:nth-child(n+3),
tr th:nth-child(n+3) {
  border-radius: 0 .625rem .625rem 0;
}

tr td:nth-child(1),
tr th:nth-child(1) {
  border-radius: .625rem 0 0 .625rem;
}
</style>