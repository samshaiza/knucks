<div id="waveform">

</div>
<script>
    export let audio_file
    import WaveSurfer from 'wavesurfer.js'

    let waveColor = "#4F4A85"
    let progressColor = "#383351"

    const player = WaveSurfer.create({
        container: '#waveform',
        waveColor: waveColor,
        progressColor: progressColor,
        url: audio_file
    })
</script>
