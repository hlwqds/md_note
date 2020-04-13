#include <stdio.h>
#include <stdint.h>

typedef struct _WAVFORMAT_
{
    char ChunkID[4];
    uint32_t ChunkSize;
    char Format[4];
    char Subchunk1ID[4];
    uint32_t Subchunk1Size;
    uint16_t AudioFormat;
    uint16_t NumChannels;
    uint32_t SampleRate;
    uint32_t ByteRate;
    uint16_t BlockAlign;
    uint16_t BitsPerSample;
    char Subchunk2ID[4];
    uint32_t Subchunk2Size;
} WAVFORMAT, *PWAVFORMAT;

int main(void)
{
    FILE *fp = fopen("test.wav", "rb");
    if (fp == NULL)
        return 0;
    WAVFORMAT wf;
    fread(&wf, sizeof(WAVFORMAT), 1, fp);
    printf("      ChunkID: %.4s\n", wf.ChunkID);
    printf("    ChunkSize: %d\n",   wf.ChunkSize);
    printf("       Format: %.4s\n", wf.Format);
    printf("  Subchunk1ID: %.4s\n", wf.Subchunk1ID);
    printf("Subchunk1Size: %d\n",   wf.Subchunk1Size);
    printf("  AudioFormat: %d\n",   wf.AudioFormat);
    printf("  NumChannels: %d\n",   wf.NumChannels);
    printf("   SampleRate: %d\n",   wf.SampleRate);
    printf("     ByteRate: %d\n",   wf.ByteRate);
    printf("   BlockAlign: %d\n",   wf.BlockAlign);
    printf("BitsPerSample: %d\n",   wf.BitsPerSample);
    printf("  Subchunk2ID: %.4s\n", wf.Subchunk2ID);
    printf("Subchunk2Size: %d\n",   wf.Subchunk2Size);
    fclose(fp);

    return 0;
}
