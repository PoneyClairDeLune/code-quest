#include <alsa/asoundlib.h>

static snd_seq_t *sequencerHandle;
static int midiInPort;

void midiOpen(void) {
	snd_seq_open(&sequencerHandle, "default", SND_SEQ_OPEN_INPUT, 0);
	snd_seq_set_client_name(sequencerHandle, "Example ALSA MIDI Listener");
	midiInPort = snd_seq_create_simple_port(sequencerHandle, "listen:in",
		SND_SEQ_PORT_CAP_WRITE | SND_SEQ_PORT_CAP_SUBS_WRITE,
		SND_SEQ_PORT_TYPE_APPLICATION
	);
}