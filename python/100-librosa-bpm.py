#!/usr/bin/env python
from __future__ import print_function
import os
import sys
import librosa
filepath = librosa.util.example_audio_file()
if len(sys.argv) == 2:
    args = sys.argv
    filepath = args[1]
filename = os.path.splitext(os.path.basename(filepath))[0]
y, sr = librosa.load(filepath)
tempo, beat_frames = librosa.beat.beat_track(y=y, sr=sr)
print('{:s}: Estimated tempo: {:.2f} BPM'.format(filename, tempo))
#beat_times = librosa.frames_to_time(beat_frames, sr=sr)
#print('Saving output to {:s}.beat_times.csv'.format(filename))
#librosa.output.times_csv('{:s}.beat_times.csv'.format(filename), beat_times)

