export class AudioRecorder {
    public audioBlobs = [] as Blob[]
    public mediaRecorder = null as MediaRecorder | null
    public streamBeingCaptured = null as MediaStream | null

    public start = () => {
        if (!(navigator.mediaDevices && navigator.mediaDevices.getUserMedia)) {
            return Promise.reject(new Error('mediaDevices API or getUserMedia method is not supported in this browser.'));
        }

        else {
            return navigator.mediaDevices.getUserMedia({ audio: true }/*of type MediaStreamConstraints*/)
                .then(stream => {
                    this.streamBeingCaptured = stream;
                    this.mediaRecorder = new MediaRecorder(stream);
                    this.audioBlobs = [];
                    this.mediaRecorder.addEventListener("dataavailable", event => {
                        this.audioBlobs.push(event.data);
                    });
                    this.mediaRecorder.start();
                });
        }
    }

    public stop = (): Promise<Blob> => {
        return new Promise(resolve => {
            const mimeType = this.mediaRecorder?.mimeType;
            this.mediaRecorder?.addEventListener("stop", () => {
                const audioBlob = new Blob(this.audioBlobs, { type: mimeType });
                resolve(audioBlob);
            });
            this.cancel();
        });
    }

    public cancel = () => {
        this.mediaRecorder?.stop();
        this.stopStream();
        this.resetRecordingProperties();
    }

    public stopStream = () => {
        this.streamBeingCaptured?.getTracks()
            .forEach(track => track.stop());
    }

    public resetRecordingProperties = () => {
        this.mediaRecorder = null;
        this.streamBeingCaptured = null;
    }
}
