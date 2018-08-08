<template lang="pug">
  .page-container
    md-app(md-waterfall md-mode="fixed")
      md-app-toolbar.md-primary
        span.md-title Web client
      md-app-content
        form.md-layout(novalidate)
          md-card.md-layout-item.md-size-50.md-small-size-100
            md-card-header
              .md-title Add serial
            md-card-content
              .md-layout.md-gutter
                .md-layout-item.md-small-size-100
                  md-field
                    label Serial name
                    md-input(name="serial")
                  div
                    md-radio(v-model="quality" value="MP4") Mp4
                    md-radio(v-model="quality" value="1080") 1080
                    md-radio(v-model="quality" value="720") 720
                  .full-control
                    md-list
                      md-subheader All serials
                      md-list-item(v-for="(serial, index) in serials", :key="index")
                        md-checkbox(v-model="serial.active")
                        span.md-list-item-text
                          | {{serial.name}} 
                          | 
                          | {{serial | prettySerial}}
                          | 
                          | {{serial.format}}

</template>

<script>
import HelloWorld from './components/HelloWorld'

export default {
  name: 'App',
  data: () => {
    return {
      quality: 'MP4',
      serials: [
        {name: "Siren", active: false, season: 7, episode: 1, format: "MP4"}
      ]
    }
  },
  filters: {
    prettySerial (serial) {
      return `(S${serial.season}E${serial.episode})`
    }
  },
  components: {}
}
</script>

<style>
  .md-drawer {
    width: 230px;
    max-width: calc(100vw - 125px);
  }
</style>
