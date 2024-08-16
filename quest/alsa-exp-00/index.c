#include "alsa/asoundlib.h"

static snd_seq_t *sequencerHandle;
static int midiInPort;

void midiOpen(void) {
	snd_seq_open(&sequencerHandle, "default", SND_SEQ_OPEN_INPUT, 0);
	snd_seq_set_client_name(sequencerHandle, "Example ALSA MIDI Node");
	midiInPort = snd_seq_create_simple_port(sequencerHandle, "i-0",
		SND_SEQ_PORT_CAP_WRITE | SND_SEQ_PORT_CAP_SUBS_WRITE,
		SND_SEQ_PORT_TYPE_APPLICATION
	);
}

snd_seq_event_t *midiRead(void) {
	snd_seq_event_t *ev = NULL;
	snd_seq_event_input(sequencerHandle, &ev);
	return ev;
}

int main() {
	midiOpen();
	while (1) {
		midiRead();
	};
	return 0;
}